package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "123.56.177.139:21")
	if err != nil {
		fmt.Println(err)
		return
	}

	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status, err)

	user, password := "ftpzy", "X1THc^enf$#o2m5L"

	conn.Write([]byte("USER " + user))
	conn.Write([]byte("PASS " + password))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Println(n)

	conn.Close()

}
