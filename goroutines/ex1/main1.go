package main

//
//var complete chan int = make(chan int)
//
//func loop1()  {
//	for i := 0; i < 10; i ++ {
//		fmt.Printf("%d ", i)
//	}
//
//	complete <- 0
//}
//
//func main() {
//	go loop1()
//	<- complete
//}

//var boos chan int = make(chan int)
//
//func boo() {
//	fmt.Println("Hello World")
//	boos <- 1
//}
//
//func main() {
//	ch := make(chan int)
//	ch <- 1 // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
//	//go boo()
//	//<- boos
//	fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
//}


//var ch1 chan int = make(chan int)
//var ch2 chan int = make(chan int)
//
//func say(s string) {
//	fmt.Println(s)
//	ch2 <- 10
//}

func main() {
	//go say("hello")
	//ch1 <- <- ch2 // ch1 等待 ch2流出的数据

	//<- ch1  // 堵塞主线

	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	//go func() {
	//	fmt.Println(123)
	//	c <- 1
	//}()
	//
	//<- c
	//time.Sleep(time.Second)
}
