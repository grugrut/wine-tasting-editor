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

	tmpl := template.Must(template.ParseFiles("../src/tmpl/index.html"))
	tmpl.Execute(w, data)
}

func handlerWinePost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	name := r.FormValue("name")
	year, _ := strconv.Atoi(r.FormValue("year"))
	created := time.Now()
	updated := time.Now()
	account := user.Current(c).String()

	wine := model.Wine{
		Name:    name,
		Year:    year,
		Account: account,
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
