package main

import (
	"errors"
	"testing"
)

func TestM(t *testing.T) {
	main()
	must(errors.New("my error"))
}
