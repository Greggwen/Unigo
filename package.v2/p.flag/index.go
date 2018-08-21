package main

import (
	"flag"
	"fmt"
)

// flag ： flag 包实现了命令行的标记解析
//func init() {
//	var ip = flag.Int("flagname", 1234, "Help message for flagname")
//
//	flag.Parse()
//	fmt.Println("ip has value ", *ip)
//}

//var flagVar string
//func init() {
//	flag.StringVar(&flagVar, "gitchat","Hello", "Help message for flagname")
//	flag.Parse()
//	fmt.Println("ip has value ", flagVar)
//}

func init() {
	// 注意： 如果是bool类型，命令行语法需要使用  -boolean=0,1,True,False,true,false,TRUE,FALSE
	// -flag=var 这种形式只支持bool类型，而-flag var 这种形式只支持非bool类型
	var ip = flag.Bool("boolean", true, "Please input true or false")
	flag.Parse()

	fmt.Println(*ip)
}

func main() {

}
