package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var IoCopyFunc = io.Copy
var IoCopyBufferFunc = io.Copy
var IoutilTempFile = ioutil.TempFile

//Mode defines the shape used when transforming images
type Mode int

//Modes supported by the primitive package.
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedrect
	ModeBeziers
	ModeRotatedellipse
	ModePolygon
)

//WithMode is an option for Transform function that will define the mode.
//Default mode is ModeTriangle
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"=m", fmt.Sprintf("%d", mode)}
	}
}

//Transform will take a provided image then apply a primitive transformation on it.
//returns a reader to the resulting image.
func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}
	in, err := tempFile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())

	out, err := tempFile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	//Read image into a file
	_, err = IoCopyFunc(in, image)
	if err != nil {
		return nil, err
	}

	//Run primitive
	stdCommand, err := primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, fmt.Errorf("primitive: Failed to run primitive command %s", stdCommand)
	}
	/*if strings.TrimSpace(stdCommand) == ""{
		panic(stdCommand)
	}
	*/
	b := bytes.NewBuffer(nil)
	_, err = IoCopyBufferFunc(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func primitive(imgIn, imgOut string, numShapes int, args ...string) (string, error) {
	argString := fmt.Sprintf("-i %s -o %s -n %d", imgIn, imgOut, numShapes)
	args = append(strings.Fields(argString), args...)
	cmd := exec.Command("primitive", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := IoutilTempFile("", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
