package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

func a() {
	defer r()
	n := []int{3, 4, 7}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}
