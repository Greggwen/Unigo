package main

import (
	"Unigo/gosftp3/components"
	"flag"
	"github.com/pkg/sftp"
	"log"
)

var (
	localBasePath  string
	remoteBasePath string
)

// TODO: 读取参数获取配置文件，后期可优化为读取配置文件，两种方案兼容
var (
	USER = flag.String("user", "xululu", "ssh username")
	HOST = flag.String("host", "192.168.1.200", "ssh server hostname")
	PORT = flag.Int("port", 22, "ssh server port")
	PASS = flag.String("pass", "mosh123", "ssh password")
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

	sftpClient, err = components.Connet(*USER, *PASS, *HOST, *PORT)

	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	components.LocalToRemote("/Users/simple.xull/DevOps/Code/local/golang/src/LeetCode-Golang", "/home/xululu/test/leetCode")

}
