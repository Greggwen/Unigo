package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type mediumStorage struct {
	mediumPath string // 媒介路径
	mediumType string // 媒介类型： file, dir
}

func main() {
	var localDirPath = "/Users/simple.xull/DevOps/Code/reworkSite/vue-framework-wz"

	var stime = time.Now()
	_, err := os.Stat(localDirPath)
	if err != nil {
		log.Fatal(err)
	}

	var medium = make(chan mediumStorage)
	var wg sync.WaitGroup

	wg.Add(1)
	go walkDir(localDirPath, &wg, medium)

	go func() {
		wg.Wait()
		close(medium)
	}()

	var twg sync.WaitGroup
	var dircount, filecount int64
	for ms := range medium {
		twg.Add(1)
		go printContent(ms, &twg)
		if ms.mediumType == "file" {
			filecount++
		} else {
			dircount++
		}
	}

	go func() {
		twg.Wait()
	}()

	var exectime = time.Since(stime)

	fmt.Printf("The total number of files is %d.\nThe total number of dir is %d\n", filecount, dircount)
	fmt.Println("The execute time is ", exectime)

}

func walkDir(dir string, wg *sync.WaitGroup, medium chan<- mediumStorage) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			medium <- mediumStorage{
				mediumPath: subdir,
				mediumType: "dir",
			}
			go walkDir(subdir, wg, medium)
		} else {
			medium <- mediumStorage{
				mediumPath: dir + "/" + entry.Name(),
				mediumType: "file",
			}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

func printContent(medium mediumStorage, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("The medium type is %s, filepath: %s\n", medium.mediumType, medium.mediumPath)
}
