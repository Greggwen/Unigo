package main

import (
	"fmt"
	"time"
)

func main() {
	// 非阻塞的收发
	noneBlocking()
	randExec()
}

// 随机执行

func randExec() {
	ch := make(chan int)

	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("case1")
		case <-ch:
			fmt.Println("case2")
		}
	}
}

// 非阻塞的收发
func noneBlocking() {
	ch := make(chan int)

	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
}
