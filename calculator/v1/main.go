package main

import "flag"

var language string

func init() {
	flag.StringVar(&language, "calculator", "en", "The language is english!")
}

func main() {
	flag.Parse()
	// 注册计算器语言包

}