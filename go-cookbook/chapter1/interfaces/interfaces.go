package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy直接将数据从in复制到out
// 它使用了buffer作为缓存
// 同时会将数据复制到Stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// 我们将结果写入 out 和 Stdout
	w := io.MultiWriter(out, os.Stdout)
	// 标准的 io.Copy 调用,如果传入的数据 in 过大会很危险
	// 第一次向w复制数据
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)
	// 使用64字节作为缓存写入 w
	// 第二次向w复制数据
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}
