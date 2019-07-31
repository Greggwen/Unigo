package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type fi struct {
	fn string
	fs int64
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- fi) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- fi{
				fn: dir + "/" + entry.Name(),
				fs: entry.Size(),
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

func main() {

	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	fileInfo := make(chan fi)

	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileInfo)
	}

	go func() {
		n.Wait()
		close(fileInfo)
	}()

	var nfiles, nbytes int64

	for fii := range fileInfo {
		nfiles++
		nbytes += fii.fs
		fmt.Println(fii.fn)
	}

	//for file := range fileName {
	//	fmt.Println(file)
	//}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
