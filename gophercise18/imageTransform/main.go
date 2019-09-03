package main

import (
	"errors"
	"fmt"
	"gophercises/gophercise18/imageTransform/primitive"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var IoutilTempFileFunc = ioutil.TempFile
var ServeAndListenFunc = http.ListenAndServe
var IoCopy = io.Copy
var StrconvAtoiFunc = strconv.Atoi

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/modify/", modifyImage)

	mux.HandleFunc("/", filePicker)

	mux.HandleFunc("/upload", uploadImage)
	fs := http.FileServer(http.Dir("./img"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	log.Fatal(ServeAndListenFunc(":5000", mux))
}

func filePicker(w http.ResponseWriter, r *http.Request) {
	html := `
	<html><body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		<input type="file" name="image">
		<button type="submit">Upload image</button>
	</form>
	</body></html>`

	fmt.Fprint(w, html)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("bad request")
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	onDiskDrive, err := tempFile("", ext)
	if err != nil {
		http.Error(w, "Something went wrong:"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer onDiskDrive.Close()
	_, err = IoCopy(onDiskDrive, file)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Println("main: Error2")
		return
	}
	http.Redirect(w, r, "/modify/"+filepath.Base(onDiskDrive.Name()), http.StatusFound)
}

func modifyImage(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	ext := filepath.Ext(f.Name())[1:]
	modeStr := r.FormValue("mode")
	if modeStr == "" {
		renderModeChoices(w, r, f, ext)
		return
	}

	mode, err := StrconvAtoiFunc(modeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nStr := r.FormValue("n")
	if nStr == "" {
		renderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}
	// numShapes, err := StrconvAtoiFunc(nStr)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// _ = numShapes
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err := generateImage(rs, ext, opt.N, opt.M)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil
}

func renderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	opts := []genOpts{
		{N: 5, M: mode},
		{N: 10, M: mode},
		{N: 15, M: mode},
		{N: 20, M: mode},
	}

	images, err := genImages(rs, ext, opts...)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	html := `<html><body>
		{{range .}}
			<a href = "/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
				<img style="width: 24%;" src="/img/{{.Name}}" />
			</a>
		{{end}}
		</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name      string
		Mode      primitive.Mode
		NumShapes int
	}
	var data []dataStruct
	for i, img := range images {
		data = append(data, dataStruct{
			Name:      filepath.Base(img),
			Mode:      opts[i].M,
			NumShapes: opts[i].N,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	opts := []genOpts{
		{N: 100, M: primitive.ModeBeziers},
		{N: 100, M: primitive.ModeCircle},
		{N: 100, M: primitive.ModeCombo},
		{N: 100, M: primitive.ModeEllipse},
	}

	images, err := genImages(rs, ext, opts...)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	html := `<html><body>
		{{range .}}
			<a href = "/modify/{{.Name}}?mode={{.Mode}}">
				<img style="width: 24%;" src="/img/{{.Name}}" />
			</a>
		{{end}}
		</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name string
		Mode primitive.Mode
	}
	var data []dataStruct
	for i, img := range images {
		data = append(data, dataStruct{
			Name: filepath.Base(img),
			Mode: opts[i].M,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func generateImage(r io.Reader, ext string, numShape int, mode primitive.Mode) (string, error) {
	out, err := primitive.Transform(r, ext, numShape, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}

	outFile, err := tempFile("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := IoutilTempFileFunc("./img/", prefix)
	if err != nil {
		return nil, errors.New("main: Failed to create temporary file.")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
