package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("hello docker golang")

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello world")
}
