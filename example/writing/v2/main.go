package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("text1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n2, "bytes written successfully")
}
