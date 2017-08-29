package server

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"io"
	"net/http"
	"net/url"
)

// Wine datamodel
type Wine struct {
	Name string
	Year int
}

type APIStatus struct {
	success bool
	code    int
	message string
}

// APIResource represents REST API Interfaces
type APIResource interface {
	Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
}

// APIResourceBase is defined for default
type APIResourceBase struct{}

// Success is the function meaning API successed
func Success(code int) APIStatus {
	return APIStatus{success: true, code: code, message: ""}
}

// Fail is the function meaning API failed
func Fail(code int, message string) APIStatus {
	return APIStatus{success: false, code: code, message: message}
}

func (APIResourceBase) Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return Fail(http.StatusMethodNotAllowed, ""), nil
}
func (APIResourceBase) Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return Fail(http.StatusMethodNotAllowed, ""), nil
}
func (APIResourceBase) Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return Fail(http.StatusMethodNotAllowed, ""), nil
}

func init() {
	http.HandleFunc("/", handlePage)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
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
