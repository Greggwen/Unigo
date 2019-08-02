package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.spingdraft.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operator timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error ", err)
		}

		return
	}
	fmt.Println(addr)
}
