package ex4

import (
	"net"
	"fmt"
	"strings"
	"errors"
	"os"
	"io"
	"strconv"
)

// 实现一个并发的FTP服务器。服务器可以解释从客户端发来的命令，例如 cd 用来改变目录，ls 用来列出目录，get 用来发送一个文件的内容，close 用来关闭连接。可以使用标准的ftp命令作为客户端，或者自己写一个。

type Ftp struct {
	con net.Conn
	ip string
}

func NewFtp(ip string) (*Ftp, error)  {
	buf := make([]byte, 1024)
	con, err := net.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}
	n, err := con.Read(buf)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(buf[:n]))
	return &Ftp{con, ip}, nil
}

func (self *Ftp) Login(user, passwd string) error {
	buf := make([]byte, 1024)
	self.con.Write([]byte(fmt.Sprintf("USER %s\r\n", user)))
	self.con.Read(buf)
	self.con.Write([]byte(fmt.Sprintf("PASSWD %s\r\n", passwd)))
	n, err := self.con.Read(buf)
	if err != nil {
		return err
	}

	if !strings.Contains(string(buf[:n]), "230 Loggend on!") {
		return errors.New(strings.TrimSpace(string(buf)))
	}

	return nil
}

func (self *Ftp) PutPasv (Pathname string) error {
	con, err := self.connections("STOR", Pathname)
	if err != nil {
		return err
	}

	File, err := os.Open(Pathname)
	if err != nil {
		con.Close()
		return err
	}

	io.Copy(con, File)
	File.Close()
	con.Close()
	buf := make([]byte, 1024)
	_, err = self.con.Read(buf)
	if err != nil {
		return err
	}
	return nil
}

func (self *Ftp) GetFile (Pathname string) error {
	con, err := self.connections("RETR", Pathname)
	if err != nil {
		return err
	}

	File, err := os.Create(Pathname)
	if err != nil {
		con.Close()
		return err
	}
	io.Copy(File, con)
	File.Close()
	con.Close()
	buf := make([]byte, 1024)
	_, err = self.con.Read(buf)
	if err != nil {
		return err
	}
	return nil
}

func (self *Ftp) connections (status, Pathname string) (net.Conn, error) {
	buf := make([]byte, 1024)
	self.con.Write([]byte("PASV \r\n"))
	n, err := self.con.Read(buf)
	if err != nil {
		return nil, err
	}

	if s := string(buf[:n]);  !strings.Contains(s, "227 Entering Passive Mode") {
		return nil, errors.New(s)
	}

	port := getport(buf[27 : n-3])
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", strings.Split(self.ip, ":")[0], port))
	if err != nil {
		return nil, err
	}
	self.con.Write([]byte(fmt.Sprintf("%s %s\r\n", status, Pathname)))
	n, err = self.con.Read(buf)
	if err != nil {
		con.Close()
		return nil, err
	}
	if !strings.Contains(string(buf[:n]), "150 Opening data channel") {
		con.Close()
		return nil, errors.New("Create data link error.")
	}
	return con, nil
}

func getport(by []byte) int {
	s := string(by)
	list := strings.Split(s, ",")
	n1, err := strconv.Atoi(list[len(list) - 2])
	if err != nil {
		return 0
	}
	n2, err := strconv.Atoi(list[len(list) - 1])
	if err != nil {
		return 0
	}
	return n1 * 256 + n2
}