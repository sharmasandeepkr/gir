package main

import (
	"labix.org/v2/mgo/bson"
)

//Service is how we serve
type Service struct {
	serviceID       string
	imageTag        string
	serviceImgL     string
	serviceImgR     string
	serviceName     string
	statitics       []string //[]filename
	statiticsP      string   //it gives comparision of all farms statitics
	datails         []string
	serviceCalendar map[string][]string
	//Details string
	//optional options
}

//User is someone registered by any 0Auth
type User struct {
	_id                       bson.ObjectId
	optedServicesForFarms     map[string][]Service //map[farm][]services
	userServiceSuggestions    map[string][]Service //map[farm][]services
	lastTenServiceSuggestions map[string][]Service //map[farm][]services

}

func main() {

}
