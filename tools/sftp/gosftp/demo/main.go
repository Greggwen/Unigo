package main

import (
	gotools "Unigo/tools/sftp/gosftp/drivers"
	"flag"
	"log"

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

	scannerDir(localDirPath)

}

func scannerDir() {

}
