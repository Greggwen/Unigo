package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	// io.Pipe()返回的r和w分别实现了io.Reader 和 io.Writer接口
	r, w := io.Pipe()
	// 在w进行写入的时候，r会发生阻塞
	go func() {
		// 这里我们简单的写入字符串，同样也可以写入json或base64编码的其他东西
		w.Write([]byte("test\n"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
