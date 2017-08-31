package server

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	"github.com/grugrut/wine-tasting-note/src/model"
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
	fmt.Fprint(w, wines)
}

func handlerWinePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
