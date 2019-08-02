package main

import "fmt"

func print(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	a := 5
	defer print(a)

	a = 100
	fmt.Println("Now value of a is ", a)
}
