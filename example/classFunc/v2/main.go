package main

import "fmt"

type add func(a int, b int) int

func main() {

	var a add = func(a int, b int) int {
		return a + b
	}

	s := a(1, 5)
	fmt.Println(s)
}
