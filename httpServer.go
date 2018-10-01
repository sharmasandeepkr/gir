package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/markbates/goth/providers/twitter"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// type Service struct {
// 	serviceID       string
// 	imageTag        string
// 	serviceImgL     string
// 	serviceImgR     string
// 	serviceName     string
// 	statitics       []string //[]filename
// 	statiticsP      string   //it gives comparision of all farms statitics
// 	datails         []string
// 	serviceCalendar map[string][]string
// 	//Details string
// 	//optional options
// }

//User is someone registered by any 0Auth
// type User struct {
// 	_id                       bson.ObjectId
// 	optedServicesForFarms     map[string][]Service //map[farm][]services
// 	userServiceSuggestions    map[string][]Service //map[farm][]services
// 	lastTenServiceSuggestions map[string][]Service //map[farm][]services
// 	oauths                    map[string]oauth
// }
// type oauth struct {
// 	name        string
// 	email       string
// 	nickname    string
// 	location    string
// 	avtarURL    string
// 	avtarimage  string
// 	description string
// 	userID      string
// 	accessToken string
// }

const (
	twitterKey    = "nX911SXh3vhMEnv9f9pMlGK6k"
	twitterSecret = "ndY07yS0KC6bGCqKmhLC1t7xkwbbTxa4n5zQGTjP8TRBAj1y9T"
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
func oauthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		http.Error(res, "not able to complete 0 authentication ", 404)
	}
	t, err := template.New("userinfo").ParseFiles("template/userInfo.html")
	if err != nil {
		http.Error(res, "not able to complete 0 authentication ", 404)
	}
	if err = t.ExecuteTemplate(res, "userInfo", user); err != nil {
		http.Error(res, "not able to complete 0 authentication ", 404)
	}

}
func signup(w http.ResponseWriter, r *http.Request) {
	t := template.New("singup")
	t = template.Must(template.ParseFiles("templates/twitterlogin.html"))
	t.ExecuteTemplate(w, "singup", nil)
}

func main() {
	goth.UseProviders(twitter.New(twitterKey, twitterSecret, "http://localhost:8080/auth/twitter/callback"))

	r := mux.NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// r.HandleFunc("/", root.ServeRoot)
	r.HandleFunc("/", temp)
	r.HandleFunc("/twitter", gothic.BeginAuthHandler)
	r.HandleFunc("/singup", signup)
	r.HandleFunc("/0auth_callback", oauthHandler)
	// r.HandleFunc("/isuggestion", mongoClient.WriteIsug)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		_ = fmt.Errorf("oh my god%v", err)
	}
	appengine.Main()
}
