package server

import (
	"net/http"
)

// HandlerUser serve page for user
func HandlerUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlerUserGet(w, r)
	case "POST":
		handlerUserPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handlerUserGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func handlerUserPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
