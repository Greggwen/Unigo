package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type mediumStorage struct {
	mediumPath string // 媒介路径
	mediumType string // 媒介类型： file, dir
}

func main() {
	var localDirPath = "/Users/simple.xull/DevOps/Code/reworkSite/vue-framework-wz"

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

	var dircount, filecount int64
	for ms := range medium {
		if ms.mediumType == "file" {
			filecount++
		} else {
			dircount++
		}
		fmt.Printf("The medium type is %s, filepath: %s\n", ms.mediumType, ms.mediumPath)
	}

	fmt.Printf("The total number of files is %d.\nThe total number of dir is %d\n", filecount, dircount)

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
