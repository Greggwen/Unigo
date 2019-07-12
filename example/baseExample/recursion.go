package main

import "fmt"

// Go语言递归函数： 递归，就是在运行的过程中调用自己
// Go语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中
// 对于解决数学上的问题是非常有用的，如计算阶乘、生成斐波那契数列等。

//func Factorial (n uint64) (result uint64) {
//	if n > 0 {
//		result = n * Factorial(n-1)
//		// 这里同样可以使用 return , 返回的参数定义为result uint64，
//		// 如果返回只是数组类型，即uint64，则必须使用 return result来返回
//		return result
//	}
//	return 1
//}
//
//
//func main() {
//	// 通过Go实现阶乘
//	var i int = 15
//	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
//}

// 斐波那契数列

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
