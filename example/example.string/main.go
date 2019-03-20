package main

import (
	"fmt"
	"strconv"
)

func main() {

	var tarsnum string = "100"

	tarsnum, err := strconv.Atoi(tarsnum)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tarsnum)
	}
}
