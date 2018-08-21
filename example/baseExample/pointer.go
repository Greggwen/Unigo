package main

import "fmt"

// Go语言指针
/**
 * 什么是指针：一个指针变量指向一个值的内存地址。
 * 声名格式： var var_name *var-type
 */

//func main() {
//	var a int = 10
//	var ip *int
//
//	ip = &a
//	fmt.Printf("a变量的地址是：%x\n", &a)
//
//	fmt.Printf("ip变量存储的地址：%x\n", ip)
//
//	fmt.Printf("*ip变量的值是%d\n", *ip)
//}

// 空指针： 当一个指针被定义后没有分配到任何一个变量时，它的值是nil
// nil指针也称为空指针，nil在概念上和其他语言的null，None，Null都一样，都指代零值或空值

// Go语言指针数组

//const MAX int = 3
//
//func main() {
//	a := []int{10, 100, 200}
//	var i int
//	var ptr [MAX]*int
//
//	for i = 0; i < MAX; i ++ {
//		ptr[i] = &a[i]
//	}
//
//	for i = 0; i < MAX; i ++ {
//		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
//	}
//
//	fmt.Println(i)
//}

// Go语言指向指针的指针

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 30001

	ptr = &a
	pptr = &ptr

	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Println(ptr)
	fmt.Printf("指向指针的指针变量 **ptr = %d\n", **pptr)
	fmt.Println(*pptr)

	fmt.Println(ptr == *pptr) // true

	// 注：ptr 与 *pptr 指向的地址是相同的
}
