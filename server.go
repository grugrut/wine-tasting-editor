package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/api/", restHandler)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Rest dummy")
}
