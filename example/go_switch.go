package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Println("Go runs on:")

// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X")
// 	case "linux":
// 		fmt.Println("Linux")
// 	default:
// 		fmt.Println("%s.", os)
// 	}
// }

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()

	// fmt.Println(today)
	// fmt.Println("today + 1 is ", today+1)
	// fmt.Println("today + 2 is ", today+2)
	// fmt.Println(time.Saturday)

	t := time.Now()

	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
}
