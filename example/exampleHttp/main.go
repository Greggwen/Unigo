package main

import (
	"Unigo/example/exampleHttp/route"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", route.RouteHandler{})

	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Println("Start http server fail:", err)
	}

}
