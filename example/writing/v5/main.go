package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func producer(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consumer(data chan int, done chan bool) {
	f, err := os.Create("concurrency.txt")
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	defer f.Close()
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			done <- false
			return
		}
	}

	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	for index := 0; index < 100; index++ {
		wg.Add(1)
		go producer(data, &wg)
	}

	go consumer(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
