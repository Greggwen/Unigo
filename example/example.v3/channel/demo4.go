package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 示例1： 只能发不能收的通道
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道
	var anotherUselessChan = make(<-chan int, 1)

	// 这里打印的是可以分别代表两个通道的指针的16进制表示
	fmt.Printf("The useless channel: %v, %v\n", uselessChan, anotherUselessChan)

	// 示例2：
	intChan1 := make(chan int, 3)
	SendIn(intChan1)

	intChan2 := getIntChan()

	// 示例4
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2:%v\n", elem)
	}

	// 示例5
	_ = GetIntChan(getIntChan)
}

func SendIn(ch chan<- int) {
	ch <- rand.Intn(1000)
}

type Notifier interface {
	SendInt(ch chan<- int)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 示例5
type GetIntChan func() <-chan int
