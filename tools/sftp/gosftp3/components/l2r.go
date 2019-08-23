package components

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// l2r: Local to remote 缩写

type fi struct {
	fn string // file name
	fs int64  // file size
}

// LocalToRemote 本地同步至远程
func LocalToRemote(localBasePath, remoteBasePath string) {
	// 遍历文件树
	fileInfo := make(chan fi)

	roots := []string{localBasePath}

	var n sync.WaitGroup

	for _, node := range roots {
		n.Add(1)
		go walkDir(node, &n, fileInfo)
	}

	go func() {
		n.Wait()
		close(fileInfo)
	}()

	var nfiles, nbytes int64

	for fii := range fileInfo {
		nfiles++
		nbytes += fii.fs
		remoteFile := filepath.Dir(remoteBasePath + "/" + fii.fn[len(localBasePath)+1:])
		//remoteFile := filepath.Dir(fii.fn[len(localBasePath) + 1:])
		//fmt.Println(remoteFile)
		parent := filepath.Dir(remoteFile)
		path := string(filepath.Separator)
		dirs := strings.Split(parent, path)
		for _, dir := range dirs {
			path = filepath.Join(path, dir)
			fmt.Println(path)
		}
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileInfo chan<- fi) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			// 忽略部分文件
			if !contaniner(filepath.Base(subdir)) {
				n.Add(1)
				go walkDir(subdir, n, fileInfo)
			}
		} else {
			fileInfo <- fi{
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

func contaniner(fact string) bool {
	ignores := []string{
		".git",
		"vendor",
	}

	for _, f := range ignores {
		if fact == f {
			return true
		}
	}

	return false
}
