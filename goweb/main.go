package main

import (
	"Unigo/goweb/controllers"
	"log"
	"net/http"
)

func main() {

	l := controllers.IndexController{}
	http.HandleFunc("/", l.Index)
	//
	log.Fatal(http.ListenAndServe(":8180", nil))
}
