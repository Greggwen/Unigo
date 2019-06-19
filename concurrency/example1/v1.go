package main

import (
	"sync"
	"fmt"
)

func main()  {
	
	test2()
	//test1()

}

func test2()  {
	var wg sync.WaitGroup

	// 打印 getter getter getter
	//for _, salutation := range []string{"hello", "welcome", "getter"} {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Println(salutation)
	//	}()
	//}

	for _, salutation := range []string{"hello", "welcome", "getter"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)

		wg.Wait()
	}

	wg.Wait()
}

func test1()  {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello")
		salutation = "welcome"
	}()
	wg.Wait()

	fmt.Println(salutation)
}
