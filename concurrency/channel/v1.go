package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

// Channel 可用于同步内存的访问，更适用的场景是goroutine 之前的传递信息
// 在Go语言中，通道是包含阻塞机制的。这意味着试图写入已满的通道的任何goroutine都会等待直到通道被清空，并且任何尝试从空闲通道读取的goroutine都将等待，直到至少有一个元素被放置
func main() {
	// t1()

	//t2()

	//t3()

	t4()
}

func t4() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Reciver %d\n", integer)
	}
}

func t3() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) //2
	wg.Wait()

}

func t2() {
	stringStream := make(chan int)
	go func() {
		defer close(stringStream)
		for i := 1; i <= 5; i++ {
			stringStream <- i
		}
	}()

	for interger := range stringStream {
		fmt.Printf("%v", interger)
	}
}

func t1() {
	// Channel 声明
	//var dataStream chan interface{}  // 1

	//dataStream := make(chan interface{})  // 2

	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello World"
	}()

	//fmt.Println(<- stringStream)
	solution, err := <-stringStream
	fmt.Printf("%v, %v", solution, err)
}
