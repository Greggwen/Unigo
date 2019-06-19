package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	fmt.Println("What's your name?")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println("Do you want make love with " + scanner.Text())
	}
}
