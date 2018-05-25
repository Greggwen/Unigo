package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println(" World!")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("Hello")
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
