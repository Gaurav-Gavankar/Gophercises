package primitive

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	f, err := os.Open("C:/go-work/src/gophercises/gophercise18/img/input/golang.png")
	if err != nil {
		fmt.Println("error while opening an image")
		return
	}
	Transform(f, "png", 10, WithMode(ModeBeziers))

	fmt.Println("testing")
}

func TestTransformNegative(t *testing.T) {
	f, err := os.Open("C:/go-work/src/gophercises/gophercise18/img/input/golang.png")
	if err != nil {
		fmt.Println("error while opening an image")
		return
	}
	_, err = Transform(f, "/", 10, WithMode(ModeBeziers))
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}

	_, err = Transform(f, " ", 10, WithMode(ModeBeziers))
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}

	tmp := IoCopyFunc
	defer func() {
		IoCopyFunc = tmp
	}()
	IoCopyFunc = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return int64(0), errors.New("Dummy error")
	}
	_, err = Transform(f, "png", 10, WithMode(ModeBeziers))
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}
}

func TestTransformNegativeAdded(t *testing.T) {
	f, err := os.Open("C:/go-work/src/gophercises/gophercise18/img/input/golang.png")
	if err != nil {
		fmt.Println("error while opening an image")
		return
	}
	tmp := IoCopyBufferFunc
	defer func() {
		IoCopyBufferFunc = tmp
	}()

	IoCopyBufferFunc = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return int64(0), errors.New("Dummy error")
	}
	_, err = Transform(f, "png", 10, WithMode(ModeBeziers))
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}
}

func TestTempFile(t *testing.T) {
	tempFile("", "")
}

func TestTempFileNegative(t *testing.T) {
	tmp := IoutilTempFile
	defer func() {
		IoutilTempFile = tmp
	}()

	IoutilTempFile = func(dir, pattern string) (f *os.File, err error) {
		return nil, errors.New("Dummy error")
	}
	_, err := tempFile("", "")
	if err == nil {
		t.Error("Expecting an error but got ", err)
	}

}

func TestPrimitive(t *testing.T) {
	primitive("testInput.jpg", "testOutput.jpg", 10, "testArg")
}
