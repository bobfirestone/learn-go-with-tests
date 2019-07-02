package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet says Hello
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler is kinda like a controller
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
