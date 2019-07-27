package transfile

import (
	gotools "Unigo/sftp/gosftp/drivers"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pkg/sftp"
)

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
	fmt.Println(*USER, *HOST, *PORT, *PASS, *SIZE)

	var (
		err        error
		sftpClient *sftp.Client
	)

	sftpClient, err = gotools.Connect(*USER, *PASS, *HOST, *PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// 用来测试的本地文件路径 和 远程机器上的文件夹
	var localFilePath = "/Users/simple.xull/DevOps/Code/local/test.php"
	var remoteDir = "/home/xululu/test/local/"
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer srcFile.Close()
	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")

}
