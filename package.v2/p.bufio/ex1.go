package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
