package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	//fmt.Println(strings.NewReader(input))
	scanner := bufio.NewScanner(strings.NewReader(input))
	//fmt.Println(scanner)
	//scanner := bufio.NewScanner(strings.NewReader(input))
	//
	//fmt.Println(scanner)

	scanner.Split(bufio.ScanWords)
	//fmt.Println(scanner)
	// count the words
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
