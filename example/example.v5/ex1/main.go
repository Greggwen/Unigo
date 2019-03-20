package main

import "fmt"

func sum(s []int, c chan int)  {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 14)

	go sum(s[:2], c)
	go sum(s[2:4], c)
	go sum(s[4:], c)

	x, y, z := <-c, <-c, <-c

	fmt.Println(x, y, z)
}


