package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alecthomas/chroma"
	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {

	tmp := ServeAndListen
	defer func() {
		ServeAndListen = tmp
	}()

	ServeAndListen = func(addr string, handler http.Handler) error {
		panic("TESTING IN PROGRESS")
	}

	assert.PanicsWithValuef(t, "TESTING IN PROGRESS", main, "they should be equal")
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
	return rr.Body.String()
}

func TestHandleDebug(t *testing.T) {
	responsestirng := CheckLinks(sourceCodeHandler, "GET", "/debug?path=C:/go-work/src/gophercises/gophercise15/recover_chroma/main.go")
	fmt.Println(responsestirng)
	// Check the response body is what we expect.
	b, err := ioutil.ReadFile("C:/go-work/src/gophercises/gophercise15/recover_chroma/main.go")
	if err != nil {
		t.Error("Reading expected file error.")
	}

	expected := string(b)

	if strings.Contains(responsestirng, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responsestirng, expected)
	} else {
		fmt.Println("Test success...")
	}
}

func TestMakeLinks(t *testing.T) {
	makeLinks("main.devMw.func1.1(0xa6cac0, 0xc000306ee0)\nC:/go-work/src/gophercises/gophercise15/recover_chroma/main.go:104 +0xaa")
	makeLinks("main.devMw.func1.1(0xa6cac0, 0xc000306ee0)\n\tC:/go-work/src/gophercises/gophercise15/recover_chroma/main.go:104 +0xaa")
}

func TestSourceCodeHandlerNegative(t *testing.T) {
	//Test: Mockeing copy function to return an Error
	tmp := FileCopyFunc
	tmp1 := GetStylesFunc

	CheckLinks(sourceCodeHandler, "GET", "/debug?line=1&path=C:/go-work/src/gophercises/gophercise15/recover_chroma/main.go")

	defer func() {
		GetStylesFunc = tmp1
	}()

	GetStylesFunc = func(name string) *chroma.Style {
		return nil
	}

	CheckLinks(sourceCodeHandler, "GET", "/debug?path=C:/go-work/src/gophercises/gophercise15/recover_chroma/main.go")

	defer func() {
		FileCopyFunc = tmp
	}()

	FileCopyFunc = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("Error while copying content from file.")
	}

	CheckLinks(sourceCodeHandler, "GET", "/debug?path=C:/go-work/src/gophercises/gophercise15/recover_chroma/main.go")

	CheckLinks(sourceCodeHandler, "GET", "/debug?path=/") //Test: wrong file path to open file

}

func TestDevMw(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeLinks(devMw(handler), "GET", "/panic")
}

func executeLinks(endpoint http.HandlerFunc, method string, url string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(endpoint)
	endpoint.ServeHTTP(rr, req)
	return rr
}
