package main

import (
	"time"
	"fmt"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func doneOver (done chan string)  {
	time.Sleep(5000 * time.Millisecond)
	done <- "over"
}

func main() {
	ch := make(chan string)
	done := make(chan string)

	go process(ch)
	fmt.Println("Step 1")
	go doneOver(done)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		case <-done:
			fmt.Println("Game Over!")
			return
		default:
			fmt.Println("no value received")
		}
	}


}
