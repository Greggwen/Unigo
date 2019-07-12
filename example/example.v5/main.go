package main

import (
	"Unigo/example/example.v5/ex4"
	"fmt"
	"log"
)

func main() {
	ftp, err := ex4.NewFtp("192.168.1.200:21")
	if err != nil {
		log.Fatal(err)
	}
	err = ftp.Login("xululu", "mosh123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ftp)
	//ftp.PutPasv("./")

	//ftp.GetFile("~/t.log")

}
