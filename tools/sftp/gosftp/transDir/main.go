package main

import (
	gotools "Unigo/tools/sftp/gosftp/drivers"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
)

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

// File status
type fileStat struct {
	name  string
	isDir bool
}

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

	var localDirPath = "/Users/simple.xull/DevOps/Code/reworkSite/swenly-note"
	var remoteDir = "/home/xululu/test/local/swenly-note"

	fileInfo, err := os.Stat(localDirPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create the remote directory if not exists
	var remotePathIsExist = true
	rfileInfo, err := sftpClient.Stat(remoteDir)
	if err != nil {
		remotePathIsExist = false
	}

	// TODO: 判断传输的本地文件是否为文件夹
	if fileInfo.IsDir() {
		// TODO: 若是文件夹， 则判断远程文件夹是否存在，不存在则创建
		if !remotePathIsExist {

			transfterDir(localDirPath, remoteDir, sftpClient)

		}

		// TODO: 传输入文件夹内容

		fmt.Println(rfileInfo)
	} else {
		// TODO: 若是文件，则判断远程路径是文件还是文件夹
		if !remotePathIsExist {
			// TODO: 判断 remoteDir 最后一个字符是否为 /
			// 若最后一个字符为 /， 则创建远程文件夹，再组成新的文件路径  remoteDir/aa.txt
			// 若最后一个字符不是/，则直接就是新的文件路径   remoteDir
		} else {
			// TODO: 判断远程路径是文件还是文件夹
			if rfileInfo.IsDir() {
				// TODO: 组成新的文件路径  remoteDir/aa.txt
			}
		}
	}

	fmt.Println(fileInfo.IsDir())

}

func transfterDir(localDirPath string, remoteDir string, sftpClient *sftp.Client) {
	// 创建文件夹
	err := sftpClient.MkdirAll(remoteDir)
	if err != nil {
		log.Fatal(err.Error())
	}

}
