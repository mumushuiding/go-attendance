package controller

import (
	"fmt"
	"net/http"
)

// HelloWorld HelloWorld
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	fmt.Fprintf(w, "hello world")
}
