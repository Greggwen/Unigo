package main

import (
	"fmt"
)

// func main() {
// 	var n [10]int /* n 是一个长度为 10 的数组 */
// 	var i, j int

// 	for i = 0; i < 10; i++ {
// 		n[i] = i + 100
// 	}

// 	for j = 0; j < 10; j++ {
// 		fmt.Printf("Element[%d] = %d\n", j, n[j])
// 	}
// }

// func main() {
// 	/* 数组 - 5 行 2 列*/
// 	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
// 	var i, j int

// 	 输出数组元素
// 	for i = 0; i < 5; i++ {
// 		for j = 0; j < 2; j++ {
// 			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
// 		}
// 	}
// }

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
