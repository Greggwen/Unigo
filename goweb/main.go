package main

import (
	"net/http"
	"Unigo/goweb/controllers"
	"log"
)


func main() {

	l := controllers.IndexController{}
	http.HandleFunc("/", l.Index)
	//
	log.Fatal(http.ListenAndServe(":8180", nil))
}
