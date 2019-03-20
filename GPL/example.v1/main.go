package main

import "fmt"

func main() {
	var p = f()
	var p1 = f()

	*p = 5
	fmt.Println(*p) // 结果：0xc4200140a0
	fmt.Println(*p1) // 结果： 0xc4200140a8

	fmt.Println(f() == f()) // false
}

// 函数返回局部变量的地址是非常安全的
func f () *int {
	var x = 1

	return &x
}
