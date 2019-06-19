package main

import (
	"fmt"
	"time"
)

func write (ch chan int)  {
	for i := 0; i < 40; i ++ {
		ch <- i
		fmt.Println("Func: successfully wrote ", i, "to ch")
	}
	close(ch)
}


func main() {
	ch := make(chan int, 4)

	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("Main read value", v,"from ch")
		//time.Sleep(2 * time.Second)
	}

}
