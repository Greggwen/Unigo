package main

import (
	"flag"
	"fmt"
)

func main() {
	// 当调用 getFlag()时，无论命令行的-chat为何值， 打印结果都为：Rolo。
	// 字符串的传递是值传递，相当于复制，当字符串改变时，是无法感知的，他们指向的不是同一个内存区域
	// 调用 getFlag1()时，打印结果就是命令的参数值，这个传递的是指针，指向同一个内存区域。
	var ip = getFlag()
	flag.Parse()
	fmt.Println(ip)
}

func getFlag() string {
	return *flag.String("chat", "Rolo", "Please Input Keyword")
}

func getFlag1() *string {
	return flag.String("chat", "Rolo", "Please Input Keyword")
}
