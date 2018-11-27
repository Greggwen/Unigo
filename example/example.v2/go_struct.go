package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Golang 指南"
	Book1.author = "小金库"
	Book1.subject = "Go语言基础"
	Book1.book_id = 1992

	Book2.title = "Python 指南"
	Book2.author = "小金库"
	Book2.subject = "Python语言基础"
	Book2.book_id = 1993

	printBook(Book1)
	printBook(Book2)

	/* 打印 Book1 信息 */
	// fmt.Printf("Book 1 title : %s\n", Book1.title)
	// fmt.Printf("Book 1 author : %s\n", Book1.author)
	// fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	// fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	// fmt.Printf("Book 2 title : %s\n", Book2.title)
	// fmt.Printf("Book 2 author : %s\n", Book2.author)
	// fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	// fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
