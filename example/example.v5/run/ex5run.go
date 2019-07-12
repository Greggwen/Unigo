package main

import (
	"Unigo/example/example.v5/ex5"
	"fmt"
	"os"
)

func main() {

	ftp := new(ex5.FTP)

	ftp.Connect("123.56.177.139", 21)
	ftp.Login("ftpzy", "X1THc^enf$#o2m5L")
	if ftp.Code == 530 {
		fmt.Println("error: login failure")
		os.Exit(-1)
	}

	ftp.List()

}
