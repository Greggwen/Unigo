package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer 会使用创建自Buffer函数的字节缓冲区
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array ❤️"

	b := Buffer(rawString)

	// 使用b.Bytes()可以快速从字节缓冲区获取字节切片
	// 使用b.String()可以快速从字节缓冲区获取字符串
	fmt.Println(b.String())

	// 由于*bytes.Buffer类型的b实现了io Reader 我们可以使用常见的reader函数
	s, err := ToString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	// 可以创建一个 bytes reader 它实现了
	// io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner
	// 接口
	reader := bytes.NewReader([]byte(rawString))
	// 我们可以使用其创建 scanner 以允许使用缓存读取和建立 token

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
