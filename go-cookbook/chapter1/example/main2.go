package main

import (
	"fmt"
	"go-cookbook/chapter1/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		fmt.Println(111)
		panic(err)
	}

	fmt.Println("Hellow")
}
