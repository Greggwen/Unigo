package main

import (
	"sync"
	"fmt"
)

var x int = 0
func increment2 (wg *sync.WaitGroup, ch chan bool)  {
	fmt.Println(x)
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 1; i <= 1000; i ++ {
		wg.Add(1)
		go increment2(&wg, ch)
	}

	wg.Wait()

	fmt.Println("final value of x", x)
}
