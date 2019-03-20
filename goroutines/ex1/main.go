package main

import (
	"fmt"
)

func loop()  {
	for i := 0; i < 10; i ++ {
		fmt.Printf("%d ", i)
	}
}

var ch chan int = make(chan int)
func foo()  {
	fmt.Println(111)
	ch <- 10
}

func main() {
	//go loop()
	//loop()
	//time.Sleep(time.Second)

	//var messages chan string = make(chan string)
	//
	//go func(message string) {
	//	messages <- message
	//}("Pong")
	//
	//fmt.Println(<-messages)


	go foo()

	<- ch
}
