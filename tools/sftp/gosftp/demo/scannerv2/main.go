package main

import (
	gotools "Unigo/tools/sftp/gosftp/drivers"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/sftp"
)

type mediumStorage struct {
	mediumPath string // 媒介路径
	mediumType string // 媒介类型： file, dir
}

var (
	// USER remote server's username
	USER = flag.String("user", "xululu", "ssh username")

	// HOST remote server's IP
	HOST = flag.String("host", "192.168.1.200", "ssh server hostname")

	// PORT remote server's ssh port
	PORT = flag.Int("port", 22, "ssh server port")

	// PASS remote server password
	PASS = flag.String("pass", "mosh123", "ssh password")

	// SIZE max packet size
	SIZE = flag.Int("s", 1<<15, "set max packet size")
)

func init() {
	flag.Parse()
}

func main() {
	var (
		err        error
		sftpClient *sftp.Client
	)

	sftpClient, err = gotools.Connect(*USER, *PASS, *HOST, *PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	var localDirPath = "/Users/simple.xull/DevOps/Code/reworkSite/vue-framework-wz"

	var stime = time.Now()
	_, err = os.Stat(localDirPath)
	if err != nil {
		log.Fatal(err)
	}

	var exectime = time.Since(stime)

	var mediumStore = make(chan mediumStorage)
	// var mediumRelease = make(chan mediumStorage)

	go scannerMedium(localDirPath, mediumStore)
	// go scannerRelease(mediumRelease, mediumStore)

	var counter int64
	var wg sync.WaitGroup
	for mr := range mediumStore {
		wg.Add(1)
		counter++
		go upload(mr, &wg, sftpClient)
	}

	go func() {
		wg.Wait()
	}()

	fmt.Println("exectime is ", exectime)
	fmt.Println("The total number of file is ", counter)
}

func upload(medium mediumStorage, wg *sync.WaitGroup, sftpClient *sftp.Client) {
	defer wg.Done()
	if medium.mediumType == "dir" {
		_, err := sftpClient.Stat(medium.mediumPath)
		if err != nil {
			err = sftpClient.MkdirAll(medium.mediumPath)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	} else {
		// srcFile, err := os.Open(medium.mediumPath)
		_, err := os.Open(medium.mediumPath)
		if err != nil {
			log.Fatal(err)
		}
		// defer srcFile.Close()
		// var remoteFileName = path.Base(medium.mediumPath)
		// dstFile, err := sftpClient.Create(path.Join(medium.mediumPath, remoteFileName))
	}
}

func scannerMedium(localpath string, mediumStore chan<- mediumStorage) {
	var wg sync.WaitGroup
	wg.Add(1)
	go walkDir(localpath, &wg, mediumStore)
	go func() {
		wg.Wait()
		close(mediumStore)
	}()
}

// func scannerRelease(mediumRelease chan<- mediumStorage, mediumStore <-chan mediumStorage) {
// 	for ms := range mediumStore {
// 		mediumRelease <- ms
// 	}
// 	close(mediumRelease)
// }

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
