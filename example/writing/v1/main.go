package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	l, err := f.WriteString("Hello Golang!")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(l, "bytes written successfully")

}
