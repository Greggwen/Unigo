package main

import (
	"sync"
	"fmt"
)

var x int = 0
func increment (wg *sync.WaitGroup)  {
	fmt.Println(x)
	x = x + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i ++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("final value of x", x)
}
