package main

import (
	"fmt"
	"net/http"
)

func service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go API!")
}

func main() {
	http.HandleFunc("/", service)
	http.ListenAndServe(":8280", nil)
}
