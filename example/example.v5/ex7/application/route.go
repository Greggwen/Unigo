package application

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		Index,
	},
	Route{
		"indexHello",
		"GET",
		"/hello",
		IndexHello,
	},
}

func Index(w http.ResponseWriter, req *http.Request)  {

}

func IndexHello(w http.ResponseWriter, req *http.Request)  {

}

func NewRoutes ()  {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.HandleFunc(route.Pattern, route.Handler)
	}

}