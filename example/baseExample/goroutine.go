package main

import "fmt"

// 示例4：通道路线实例
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

// 示例3：通道同步实例
//func worker(done chan bool) {
//	fmt.Println("Working...")
//	time.Sleep(time.Second)
//	fmt.Println("done")
//
//	done <- true
//}
//
//func main() {
//	done := make(chan bool, 1)
//	go worker(done)
//
//	//<-done
//	fmt.Println(<-done)
//}

// 示例2 通道缓冲实例
//func main() {
//	messages := make(chan string, 2)
//
//	messages <- "Hello"
//	messages <- "World"
//
//	fmt.Println(<-messages)
//	fmt.Println(<-messages)
//}

// 示例1  通道实例
//func main() {
//	message := make(chan string)
//
//	go func() {message <- "ping"}()
//
//	msg := <-message
//
//	fmt.Println(msg)
//}
