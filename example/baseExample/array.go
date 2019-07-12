package main

import "fmt"

// 声明数组： var variable_name [SIZE] variable_type

//func main() {
//	// 初始化数组
//	//var balance = [5]float32{1000.0, 2.0, 3.4, 4.2, 4}
//	//fmt.Println(balance)
//
//	var n [10] int
//	var i, j int
//
//	for i = 0; i < 10; i ++ {
//		n[i] = i + 100
//	}
//
//	for j = 0; j < 10; j ++ {
//		fmt.Printf("Element[%d] = %d\n", j, n[j])
//	}
//}

// 多维数组
//func main() {
//	var a = [5][2]int{{0, 0}, {1, 2}, {2, 3}, {4, 5}, {5, 3}}
//	var i, j int
//
//	for i = 0; i < 5; i++ {
//		for j = 0; j < 2; j++ {
//			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
//		}
//	}
//}

// 函数传递数组
func main() {
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance, 5)

	fmt.Printf("平均值为：%f", avg)
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
