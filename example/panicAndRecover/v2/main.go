package main

import (
	"fmt"
	"time"
)

func main() {
	a()
	fmt.Println("normally returned from main")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered ", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")

	b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("ho! B paniced")
}
