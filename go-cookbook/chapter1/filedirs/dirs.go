package filedirs

import "os"

// 操作文件和目录
func Operate() error {
	if err := os.Mkdir("./jojo", os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.Chdir("../"); err != nil {
		return err
	}

	return nil
}
