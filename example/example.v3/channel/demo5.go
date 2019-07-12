package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example1() {
	// 准备好几个通道
	intChannel := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	// 随机选择一个通道，并向它发送元素值
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannel[index] <- index

	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行
	select {
	case <-intChannel[0]:
		fmt.Println("The first candidate case is selected")
	case <-intChannel[1]:
		fmt.Println("THe second candidate case is selected")
	case elem := <-intChannel[2]:
		fmt.Println("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected")

	}

}

// 示例2。
func example2() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

func main() {
	example2()
}
