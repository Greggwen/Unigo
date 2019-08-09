package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("append.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("file appended successfully")
}
