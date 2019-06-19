package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type fi struct {
	name string
	time time.Time
}

func walkDir(dir string, fileName chan<- fi) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileName)
		} else {

			info := fi{
				dir + entry.Name(),
				entry.ModTime(),
			}

			fileName <- info
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		log.Fatal("No dist directory")
		os.Exit(1)
	}

	// fileName := make(chan string)
	fileName := make(chan fi)

	go func() {
		for _, root := range roots {
			walkDir(root, fileName)
		}

		close(fileName)
	}()

	today := time.Now()
	// fmt.Println(today.Date())
	// td := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", today.Year(), today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second())

	// td := today.Format("2006-01-02")
	today = today.AddDate(0, 0, -3)
	td := today.Format("2006-01-02")
	for file := range fileName {
		time := file.time.Format("2006-01-02")
		if time < td {
			err := os.Remove(file.name)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("The file can't remove!")
		}
	}

}
