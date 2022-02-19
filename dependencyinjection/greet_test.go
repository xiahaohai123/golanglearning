package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	var buffer = bytes.Buffer{}
	buffer.Write(nil)
	//Greet(buffer, "Chris")
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	Greet(os.Stdout, "laizj")
}
