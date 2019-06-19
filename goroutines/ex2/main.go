package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Game Over")
				return
			default:
				fmt.Println("进行中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func chanSelectDemo() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("已退出，再见")
				return
			default:
				fmt.Println("goruntine 进行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("要不要停止？")
	stop <- true
	time.Sleep(5 * time.Second)
}

func wgDemo() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Wait()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Wait()
	}()

	wg.Done()

	fmt.Println("好了，大家都干完了，放工")
}
