package main

import "fmt"

type Printer func(content string) (n int, err error)

// 什么是高阶函数？ 满足以下其中任意一个就可以称为高阶函数
// 1. 接受其他的函数作为参数传入
// 2. 把其他的函数作为结果返回

// 可以看作函数 PrintToStd 是 Printer 的一个实现
func PrintToStd(content string) (bytesNum int, err error) {

	return fmt.Println(content)
}

func main() {
	var p Printer
	p = PrintToStd
	p("Hello World")
}
