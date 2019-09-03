package main

import (
	"bytes"
	"errors"
	"fmt"
	"gophercises/gophercise18/imageTransform/primitive"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tmp := ServeAndListenFunc
	defer func() {
		ServeAndListenFunc = tmp
	}()

	ServeAndListenFunc = func(addr string, handler http.Handler) error {
		panic("TESTING IN PROGRESS")
	}

	assert.PanicsWithValuef(t, "TESTING IN PROGRESS", main, "they should be equal")

}

func TestGenImages(t *testing.T) {
	opts := []genOpts{
		{N: 100, M: primitive.ModeEllipse},
		{N: 200, M: primitive.ModeRect},
	}

	ext := "jpeg"
	var rs, err = os.Open("C:/go-work/src/gophercises/gophercise18/img/input/golang.png")
	_, err = genImages(rs, ext, opts...)
	if err != nil {
		t.Error("Error while generating image.")
	}
}

type fn func(resp http.ResponseWriter, rq *http.Request)

func CheckLinks(endpoint fn, method string, url string) string {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(endpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	return rr.Body.String()
}

func TestGenerateImage(t *testing.T) {
	file, err := os.Open("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	generateImage(file, "*", 1, primitive.ModeBeziers)
}

func TestFilePicker(t *testing.T) {
	CheckLinks(filePicker, "GET", "/")
}

func TestUploadImage(t *testing.T) {
	file, _ := os.Open("./img/testImages/971058115.jpg")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", "test.jpg")
	io.Copy(part, file)
	writer.Close()
	r, _ := http.NewRequest("POST", "localhost:9000/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fn(uploadImage))
	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusOK {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//mocking io.Copy function to return an error
	tmp := IoCopy
	defer func() {
		IoCopy = tmp
	}()

	IoCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("TEST ERROR")
	}
	handler.ServeHTTP(rr, r)
}

func TestUploadImageNegative(t *testing.T) {
	file, _ := os.Open("./img/testImages/971058115.png")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", "test.jpg")
	io.Copy(part, file)
	writer.Close()
	r, _ := http.NewRequest("POST", "localhost:9000/upload", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fn(uploadImage))
	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusOK {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	part, _ = writer.CreateFormFile("image", "test./")
	io.Copy(part, file)
	writer.Close()
	// r, _ = http.NewRequest("POST", "localhost:9000/upload", body)
	// r.Header.Set("Content-Type", writer.FormDataContentType())

	// rr = httptest.NewRecorder()
	// handler = http.HandlerFunc(fn(uploadImage))
	handler.ServeHTTP(rr, r)

}

//TO generate error for creating tempfile
func TestUploadImageNegativeAdded(t *testing.T) {
	file, _ := os.Open("./img/testImages/971058115.jpg")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", "test.*")
	io.Copy(part, file)
	writer.Close()
	r, _ := http.NewRequest("POST", "localhost:9000/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fn(uploadImage))
	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusOK {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	handler.ServeHTTP(rr, r)
}

func TestModifyImage(t *testing.T) {
	urls := []string{
		"http://localhost:9000/modify/087884137.png?mode=3",
		"http://localhost:9000/modify/407308789.png?mode=3&n=400",
		"http://localhost:9000/modify/087884137.png",
		"http://localhost:9000/modify/",
	}

	for _, url := range urls {
		CheckLinks(modifyImage, "POST", url)
	}
}

func TestModifyImageNegative(t *testing.T) {
	tmp := StrconvAtoiFunc
	CheckLinks(modifyImage, "POST", "http://localhost:9000/modify/407308789.png/?mode=*")
	StrconvAtoiFunc = tmp
}
func TestTempFile(t *testing.T) {
	tmp := IoutilTempFileFunc
	defer func() {
		IoutilTempFileFunc = tmp
	}()
	IoutilTempFileFunc = func(dir, pattern string) (f *os.File, err error) {
		return nil, errors.New("TEST ERROR")
	}
	tempFile("test", "jpg")
}

func TestGenImage(t *testing.T) {
	opts := []genOpts{
		{N: 5, M: primitive.ModeBeziers},
	}
	file, _ := os.Open("./img/testImages/971058115.png")
	genImages(file, "*", opts...)
}

func TestRenderNumShapeChoices(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9000/modify/407308789.png", nil)
	if err != nil {
		fmt.Println(err)
	}
	rr := httptest.NewRecorder()

	file, err := os.Open("./img/testImages/971058115.png")
	if err != nil {
		fmt.Println(err)
	}
	renderNumShapeChoices(rr, req, file, "*", primitive.ModeBeziers)
}

func TestRenderModeChoices(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9000/modify/407308789.png", nil)
	if err != nil {
		fmt.Println(err)
	}
	rw := httptest.NewRecorder()

	file, err := os.Open("./img/testImages/971058115.png")
	if err != nil {
		fmt.Println(err)
	}
	renderModeChoices(rw, req, file, "*")
}
