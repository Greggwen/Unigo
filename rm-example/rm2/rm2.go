package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type singleFile struct {
	name string
	time time.Time
}

func walkDir(dir string, n *sync.WaitGroup, needDelFiles chan<- singleFile) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir+"/", n, needDelFiles)
		} else {
			needDelFiles <- singleFile{
				dir + entry.Name(),
				entry.ModTime(),
			}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entiries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entiries
}

func main() {

	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		log.Fatal("No dist directory")
		os.Exit(1)
	}

	needDelFiles := make(chan singleFile)

	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, needDelFiles)
	}

	go func() {
		n.Wait()
		close(needDelFiles)
	}()

	today := time.Now()

	today = today.AddDate(0, 0, -3)
	td := today.Format("2006-01-02")

	for file := range needDelFiles {
		time := file.time.Format("2006-01-02")
		if time < td {
			// err := os.Remove(file.name)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			log.Println("Remove the file path of " + file.name)
		}
	}

}
