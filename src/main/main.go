package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There! This seems great1 %s", r)
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8888", nil)
}
