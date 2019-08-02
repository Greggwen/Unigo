package main

import "fmt"

func main() {
	str := "Hello World"
	fmt.Printf("Origined String is %s\n", str)
	fmt.Printf("Reversed String is ")

	for _, v := range []rune(str) {
		defer fmt.Printf("%c", v)
	}
}
