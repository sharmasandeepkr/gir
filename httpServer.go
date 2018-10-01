package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["home"] = template.Must(template.ParseFiles("templates/nav.html", "templates/farmSmallDevices.html", "templates/optedServices.html", "templates/suggestions.html", "templates/statitics.html", "templates/farmLargeDevices.html", "templates/home.html"))
}
func temp(w http.ResponseWriter, r *http.Request) {
	err := templates["home"].ExecuteTemplate(w, "home", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// r.HandleFunc("/", root.ServeRoot)
	r.HandleFunc("/", temp)
	// r.HandleFunc("/isuggestion", mongoClient.WriteIsug)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_ = fmt.Errorf("oh my god%v", err)
	}
	appengine.Main()
}
