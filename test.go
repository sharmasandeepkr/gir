package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"structData"
)

// func init() {
// 	if err := mongoClient.MyDB.ConnectDB(); err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(3)
// 	}
// }
var T *template.Template

func init() {
	// var templates map[string]string
	// templates = make(map[string]string)
	var err error
	// T, err = template.ParseFiles("sin_templates/test.html")
	// T, err = template.ParseFiles("sin_templates/sin_farmLargeDevices.html")
	// T, err = template.ParseFiles("sin_templates/sin_nav.html")
	// T, err = template.ParseFiles("sin_templates/sin_statistics.html")
	// T, err = template.ParseFiles("sin_templates/sin_optedServices.html")
	// T, err = template.ParseFiles("sin_templates/sin_suggestions.html")
	// T, err = template.ParseFiles("sin_templates/sin_farmSmallDevices.html")
	T, err = template.ParseFiles("sin_templates/sin_home.html", "sin_templates/sin_nav.html", "sin_templates/sin_farmSmallDevices.html", "sin_templates/sin_optedServices.html", "sin_templates/sin_suggestions.html", "sin_templates/sin_statistics.html", "sin_templates/sin_farmLargeDevices.html")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(999)
	}
}

// type ldata struct {
// 	Ucred            `bson:"ucred"`
// 	FarmLargeDevices `bson:"farmLargeDevices"`
// }
// type Ucred struct {
// 	Picture string `bson:"picture"`
// }
// type FarmLargeDevices struct {
// 	// Farms map[string]string `bson:"farms"`
// 	Adv `bson:"adv"`
// }
// type Adv struct {
// 	ImgURL  string `bson:"imgurl"`
// 	Heading string `bson:"heading"`
// }

func navHandler(w http.ResponseWriter, r *http.Request) {

	err := T.ExecuteTemplate(w, "sin_home", &structData.SinHomeInstance)
	_ = T
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// defer mongoClient.MyDB.Session.Close()

	http.Handle("/", http.HandlerFunc(navHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("so series")
	}
}
