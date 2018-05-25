package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(3, 5))
	fmt.Println(add1(3, 5))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
