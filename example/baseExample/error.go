package main

import "fmt"

// Go语言通过内置的错误接口提供非常简单的错误处理机制
// error 类型是一个接口类型，如下
// type error interface {
//	Error() string
//}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			varDividee,
			varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is:", errorMsg)
	}
}
