package server

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	"github.com/grugrut/wine-tasting-note/src/model"
	"strconv"
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
	html := `
<html>
<body>
<form method="post">
<input name="name">
<input name="year">
<input type="submit">
</form>
</body>
</html>
`
	fmt.Fprintln(w, html)
	fmt.Fprintln(w, wines)
}

func handlerWinePost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	year, _ := strconv.Atoi(r.FormValue("year"))
	wine := model.Wine{
		Name: r.FormValue("name"),
		Year: year,
	}
	err := model.InsertWine(c, wine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/wine", http.StatusFound)
}
