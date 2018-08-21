package main

import "fmt"

// Go语言结构体： 结构体是由一系列具有相同类型或不同类型的数组构成的数据集合
// 定义结构体： 结构体定义需要使用type和struct语句。struct语句定义一个新的数据类型，结构体中有一个或多个成员。type语句设定了结构体的名称。

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

//func main() {
//	// 创建一个新的结构体
//	fmt.Println(Books{"go 语言", "www.baidu.com", "Go 语言教程", 7229})
//
//	// 使用 key => value格式创建
//	fmt.Println(Books{title: "Go 语言", author: "www.baidu.com", subject: "Go 语言教程", book_id: 7229})
//
//	// 忽略字段为 0 或空
//	fmt.Println(Books{title: "Go 语言", subject: "Go 语言教程"})
//}

// 访问结构体成员

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "小金库so"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 7229

	Book2.title = "Python 语言"
	Book2.author = "小金库so"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 7228

	//fmt.Printf("Book 1 title: %s\n", Book1.title)
	//fmt.Printf("Book 1 author: %s\n", Book1.author)
	//fmt.Printf("Book 1 subject: %s\n", Book1.subject)
	//fmt.Printf("Book 1 book_id: %d\n", Book1.book_id)
	//
	//fmt.Printf("Book 2 title: %s\n", Book2.title)
	//fmt.Printf("Book 2 author: %s\n", Book2.author)
	//fmt.Printf("Book 2 subject: %s\n", Book2.subject)
	//fmt.Printf("Book 2 book_id: %d\n", Book2.book_id)

	printBook(&Book1)
	printBook(&Book2)

}

// 结构体作为函数参数
//func printBook(book Books) {
//	fmt.Printf("Book title: %s\n", book.title)
//	fmt.Printf("Book author: %s\n", book.author)
//	fmt.Printf("Book subject: %s\n", book.subject)
//	fmt.Printf("Book book_id: %d\n", book.book_id)
//}

// 结构体指针
func printBook(book *Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)

	fmt.Println(&book.title)
	fmt.Println(&book.author)
	fmt.Println(&book.subject)
	fmt.Println(&book.book_id)
}
