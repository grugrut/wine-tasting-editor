package server

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"net/http"
)

// Wine datamodel
type Wine struct {
	Name string
	Year int
}

func init() {
	http.HandleFunc("/", pageHandler)
	http.HandlerFunc("/wine", wineHandler)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if u := user.Current(c); u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	q := datastore.NewQuery("Wine").Order("-Name").Limit(10)
	wines := make([]Wine, 0, 10)
	if _, err := q.GetAll(c, &wines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, wines)
}

func wineHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if u := user.Current(c); u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	switch r.Method {
	case "GET":
		//wineGetHandler(w, r)
		return
	case "POST":
		//winePostHandler(w, r)
		return
	case "DELETE":
		//wineDeleteHandler(w, r)
		return
	}
	http.Error(w, "NotFound", http.StatusNotFound)
}
