package server

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	"github.com/grugrut/wine-tasting-note/src/model"
	"google.golang.org/appengine/user"
	"html/template"
	"strconv"
	"time"
)

// HandlerWine serve page for wine
func HandlerWine(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlerWineGet(w, r)
	case "POST":
		handlerWinePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handlerWineGet(w http.ResponseWriter, r *http.Request) {
	wines, err := model.GetWines(appengine.NewContext(r))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	data := map[string]interface{}{
		"Wines": wines,
	}

	tmpl := template.Must(template.ParseFiles("../src/tmpl/base.html", "../src/tmpl/wine.html"))
	tmpl.Execute(w, data)
}

func handlerWinePost(w http.ResponseWriter, r *http.Request) {
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
	name := r.FormValue("name")
	vintage, _ := strconv.Atoi(r.FormValue("vintage"))
	who := r.FormValue("who")
	where := r.FormValue("where")
	winetype := r.FormValue("type")
	wineurl := r.FormValue("url")
	created := time.Now()
	updated := time.Now()

	wine := model.Wine{
		Name:    name,
		Vintage: vintage,
		Who:     who,
		Where:   where,
		Type:    winetype,
		URL:     wineurl,
		Created: created,
		Updated: updated,
	}
	err := model.InsertWine(c, wine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/wine", http.StatusFound)
}
