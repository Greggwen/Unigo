package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// bytes 和 strings 包
func Buffer(rawString string) *bytes.Buffer {
	// 将传入的字符串转换为字节数组
	rawBytes := []byte(rawString)
	// 有很多方式使用字节数组或原始字符串建立缓冲区
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// 或者
	//b = bytes.NewBuffer(rawBytes)

	// 使用字符串建立字节数组
	//b = bytes.NewBufferString(rawString)

	return b
}

// ToString 接受 io.Reader, 并将其转换为字符串返回
func ToString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
