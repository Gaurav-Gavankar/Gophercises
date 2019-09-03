package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	main()
	must(errors.New("my error"))
}
