package main

// waitGroup
import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done() // 减少计数
		fmt.Printf("Hello form %v\n", id)
	}

	const numGreeters = 5

	var wg sync.WaitGroup
	for i := 0; i <= numGreeters; i++ {
		wg.Add(1) // 增加计数
		go hello(&wg, i)
	}
	wg.Wait() // 调用Wait会阻塞并等待至计数器归零。

}
