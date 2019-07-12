package main

import "fmt"

// Go语言变量作用域
/**
 * 函数内定义的变量称为局部变量
 * 函数外定义的变量称为全局变量
 * 函数定义中的变量称为形式参数
 */

// 局部变量
//func main() {
//	var a, b, c int
//	a = 10
//	b = 20
//	c = a + b
//	fmt.Printf("结果：a = %d, b = %d, c = %d\n", a, b ,c)
//}

// 全局变量
// Go语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。（就近原则）
//var g int = 20
//func main() {
//	//var a, b int
//	//a = 10
//	//b = 20
//	//g = a + b
//	//fmt.Printf("结果：a = %d, b = %d, g = %d\n", a, b ,g)
//
//	var g int = 10
//	fmt.Printf("结果： g = %d\n", g) // 输出： 结果： g = 10
//}

// 形式参数
var a int = 20

func main() {
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中的a=%d\n", a)
	c = sum(a, b)

	fmt.Printf("main()函数 c= %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum()函数中的a = %d\n", a)
	fmt.Printf("sum()函数中的b = %d\n", b)
	return a + b
}

// 初始化局部和全局变量的默认值分别为：
// int 默认值为0   float32 默认值为0  pointer默认值为nil
