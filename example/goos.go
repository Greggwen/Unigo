package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Println("The operate system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Println("The path is: %s\n", path)
}
