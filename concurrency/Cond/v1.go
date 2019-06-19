package main

import (
	"sync"
	"time"
	"fmt"
)

// Cond 实现了一个条件变量，用于等待或宣布事件发生时goroutine的交汇点
func main () {

	c := sync.NewCond(&sync.Mutex{})  // 使用一个标准的sync.Mutex作为Locker来创建Cond
	queue := make([]interface{}, 0, 10) // 创建一个长度为零的切片。 由于我们知道最终会添加10个元素，因此我们将其容量设为10。

	removeFromQueue := func (delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()  // 再次进入该并发条件下的关键部分，以修改与并发条件判断直接相关的数据
		queue = queue[1:] // 移除切片的头部并重新分配给第二个元素，这一步模拟了元素出列
		fmt.Println("Remove from queue")
		c.L.Unlock()  // 退出操作关键部分，因为我们已经成功移除了一个元素
		c.Signal()  // 发出信号，通知处于等待状态的goroutine可以进行下一步了
	}

	for i := 0; i < 10; i ++ {
		c.L.Lock()  // 在进入关键的部分前调用Lock来锁定c.L
		for len(queue) == 3 {// 在这里我们检查队列的长度，以确认什么时候需要等待。由于removeFromQueue是异步的，for不满足时才会跳出，而if做不到重复判断，这一点很重要
			c.Wait()  // 调用Wait，这将阻塞main goroutine，直到接受到信号
		}
		fmt.Println("Adding to Queue")
		queue = append(queue, struct {}{})
		go removeFromQueue(1 * time.Second)  // 创建一个新的goroutine，它会在1秒后将元素移出队列
		c.L.Unlock() // 退出条件的关键部分，因为我们已经成功加入了一个元素
	}

	time.Sleep(2 * time.Second)
}
