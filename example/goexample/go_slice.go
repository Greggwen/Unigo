package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%v\n", slice1)

	slice2 := slice1[3:8]
	fmt.Printf("%v\n", slice2)

	slice1[3:8][3] = 18
	fmt.Printf("%v\n", slice2)
	fmt.Printf("%v\n", slice1)

}
