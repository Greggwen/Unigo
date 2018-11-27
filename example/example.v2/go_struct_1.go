package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// data := make(map[string]interface{}, 4)

	// data["title"] = "Golang 指南"
	// data["author"] = "小金库"
	// data["subject"] = "Go语言基础"
	// data["book_id"] = 1992

	var Book1 Books

	Book1.title = "Golang 指南"
	Book1.author = "小金库"
	Book1.subject = "Go语言基础"
	Book1.book_id = 1992

	info := *Book1.(main.Books)

	fmt.Println(info)
	// fmt.Println(ok)
}
