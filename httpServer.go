package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"signin"
	"sinit"
	"structData"

	"github.com/gorilla/mux"
)

var SinT *template.Template

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

// var templates map[string]*template.Template

func init() {
	sinit.Init()
	// err := mongoClient.MyDB.ConnectDB()
	// if err != nil {
	// 	os.Exit(132)
	// }
	var err error
	SinT, err = template.ParseFiles("sin_templates/sin_home.html", "sin_templates/sin_nav.html", "sin_templates/sin_farmSmallDevices.html", "sin_templates/sin_optedServices.html", "sin_templates/sin_suggestions.html", "sin_templates/sin_statistics.html", "sin_templates/sin_farmLargeDevices.html", "sin_templates/sin_Products.html")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(999)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// err = sinit.Templates["home"].ExecuteTemplate(w, "home", &genHomeData)
	// if err != nil {
	// 	os.Exit(111)
	// 	//http.Error(w, "execution fails", 401)
	// }
	err := SinT.ExecuteTemplate(w, "sin_home", &structData.SinHomeInstance)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	}
}
func hsin(w http.ResponseWriter, r *http.Request) {
	var sinData structData.SinHome
	sinData = structData.SinHomeInstance
	err := SinT.ExecuteTemplate(w, "sin_home", &sinData)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	}

}

func decide(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

func main() {
	fmt.Println("i am starting with all due permission to Gau Maa!")
	// defer mongoClient.MyDB.Session.Close()
	// h := adapt(http.HandlerFunc(homeHandler), withDB(mongoClient.MyDB))
	r := mux.NewRouter().StrictSlash(true)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/auth/gplus", signin.HandleLogin)
	r.Handle("/auth/gplus/callback", signin.HandlerCallback(http.HandlerFunc(signin.HandleCallback)))
	r.HandleFunc("/sin", hsin)
	// r.Handle("/home", context.ClearHandler(h))
	r.HandleFunc("/home", homeHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Print(fmt.Errorf("oh my god%v", err))
		os.Exit(11)
	}

}
