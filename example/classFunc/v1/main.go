package main

import (
	"fmt"
)

func main() {
	// a := func() {
	// 	fmt.Println("Hello world first class function")
	// }

	// fmt.Printf("%T\n", a)

	// func() {
	// 	fmt.Println("Hello world first class function 1!")
	// }()

	func(name string) {
		fmt.Println("Hello world first class function 1!", name)
	}("Gopher")
}
