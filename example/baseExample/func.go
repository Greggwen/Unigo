package main

import "fmt"

// Go语言函数
func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a的值是:%d\n", a)
	fmt.Printf("交换前，b的值是:%d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后，a的值是:%d\n", a)
	fmt.Printf("交换后，b的值是:%d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
