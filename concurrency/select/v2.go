package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var vCount1, vCount2 int
	for i := 0; i < 1000; i ++ {
		select {
		case <-c1:
			vCount1 ++
		case <-c2:
			vCount2 ++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", vCount1, vCount2)
}
