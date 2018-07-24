package main

import "fmt"

func main() {
	var n [10]int

	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var balance = []int{1000, 2, 3, 17, 50}

	var avg float32

	avg = getAvage(balance, 5)

	fmt.Printf("平均值为：%f", avg)
	fmt.Println(balance)
}

func getAvage(arr []int, size int) float32 {

	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
		arr[i] = arr[i] + 100
	}

	avg = float32(sum) / float32(size)

	return avg
}
