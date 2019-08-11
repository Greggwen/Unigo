package main

// sync.mutex 解决并发竞争条件的问题
import (
	"sync"
	"fmt"
)

var x int = 0
func increment1 (wg *sync.WaitGroup, mu *sync.Mutex)  {
	fmt.Println(x)
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 1000; i ++ {
		wg.Add(1)
		go increment1(&wg, &mu)
	}

	wg.Wait()

	fmt.Println("final value of x", x)
}
