package server

import (
	"net/http"

	"github.com/grugrut/wine-tasting-note/src/server"
)

func init() {
	http.HandleFunc("/", handlePage)
	http.HandleFunc("/wine/", server.HandlerWine)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	/*	c := appengine.NewContext(r)
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
		fmt.Fprint(w, wines)*/
}
