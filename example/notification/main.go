package main

import "fmt"

func get_notification(user string) chan string {
	notification := make(chan string)

	go func() {
		notification <- fmt.Sprintf("Hi %s, welcome to here", user)
	}()

	return notification
}

func main() {
	jack := get_notification("jack")
	joe := get_notification("joe")

	fmt.Println(<-jack)
	fmt.Println(<-joe)
}
