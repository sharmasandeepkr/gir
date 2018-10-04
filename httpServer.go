package main

import (
	"fmt"
	"html/template"
	"mongoClient"
	"net/http"
	"signin"
	"sinit"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// type adapter func(http.Handler) http.Handler

// func adapt(h http.Handler, adpaters ...adapter) http.Handler {
// 	for _, adapter := range adpaters {
// 		h = adapter(h)
// 	}
// 	return h
// }

// func withDB(db mongoClient.DataStore) adapter {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			dbSession := db.Copy()
// 			defer dbSession.Close()
// 			context.Set(r, "database", dbSession)
// 			h.ServeHTTP(w, r)

// 		})
// 	}
// }

var templates map[string]*template.Template

func init() {
	sinit.Init()
	mongoClient.MyDB.ConnectDB()

}

func temp(w http.ResponseWriter, r *http.Request) {
	ses := mongoClient.MyDB.Session
	myhandlerSes := ses
	// defer myhandlerSes.Close()
	col := myhandlerSes.DB("girgoras").C("isuggestions")
	err := sinit.Templates["home"].ExecuteTemplate(w, "home", col)
	if err != nil {
		http.Error(w, "execution fails", 401)
	}
}
func hsin(w http.ResponseWriter, r *http.Request) {
	email := w.Header().Get("ggid")
	if email == "" {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	}
	err := sinit.Templates["sin"].ExecuteTemplate(w, "sin_home", email)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	}

}

func decide(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func main() {
	fmt.Println("i am starting with all due permission to Gau Maa!")
	defer mongoClient.MyDB.Close()
	// h := adapt(http.HandlerFunc(temp), withDB(mongoClient.MyDB))
	r := mux.NewRouter().StrictSlash(true)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Handle("/", signin.Hsin(http.HandlerFunc(decide)))
	r.HandleFunc("/auth/gplus", signin.HandleLogin)
	r.Handle("/auth/gplus/callback", signin.HandlerCallback(http.HandlerFunc(signin.HandleCallback)))
	r.HandleFunc("/sin", hsin)
	// r.Handle("/home", context.ClearHandler(h))
	r.HandleFunc("/home", temp)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_ = fmt.Errorf("oh my god%v", err)
	}
	appengine.Main()

}
