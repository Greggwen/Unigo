package main

import (
	"Unigo/example/example.v5/ex7/application"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func main() {

	application.NewRoutes()

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)
	//mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	io.WriteString(writer, "Hello hello")
	//})
	//err := http.ListenAndServe(":8081", mux)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
