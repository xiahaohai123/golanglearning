package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

type AAA func(w http.ResponseWriter, r *http.Request)

func (receiver AAA) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	_ = http.ListenAndServe(":5000", AAA(MyGreeterHandler))
}
