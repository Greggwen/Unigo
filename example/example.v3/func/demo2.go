package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 加
func Sum (x, y int) int {
	return x + y
}

// 减
func Reduce (x, y int) int {
	return x - y
}

// 乘
func Pursue(x, y int) int {
	return x * y
}

// 除
func Division (x, y int) int {
	return x / y
}

func calculate (x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}

	return op(x, y), nil
}

func main() {
	var op operate
	op = Division
	fmt.Println(op(3, 5))
}
