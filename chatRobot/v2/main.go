package main

import (
	"Unigo/chatRobot/v2/chatbot"
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

// chatbotName 决定了对话使用的聊天机器人
var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.en", "The chatbot's name for dialogue.")
}

func main() {
	flag.Parse()

	// 注册聊天机器人
	chatbot.Register(chatbot.NewSimpleEN("simple.en", nil))
	chatbot.Register(chatbot.NewSimpleCn("simple.cn", nil))

	// 根据选择获取机器人，默认为simple.en
	myChatbot := chatbot.Get(chatbotName)

	// 如果chatbot 不存在
	if myChatbot == nil {
		err := fmt.Errorf("Fatal error: Unsupported chatbot named %s\n", chatbotName)
		checkError(nil, err, true)
	}

	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	// 开始聊天了
	begin, err := myChatbot.Begin()
	// 检测err，如果err不为nil，就退出
	checkError(myChatbot, err, true)
	fmt.Println(begin)

	// 读取客户端的输入
	input, err := inputReader.ReadString('\n')
	checkError(myChatbot, err, true)
	fmt.Println(myChatbot.Hello(input[:len(input)-1]))

	for {
		input, err = inputReader.ReadString('\n')
		if checkError(myChatbot, err, false) {
			continue
		}

		output, end, err := myChatbot.Talk(input)

		if checkError(myChatbot, err, false) {
			continue
		}

		if output != "" {
			fmt.Println(output)
		}

		if end {
			err = myChatbot.End()
			checkError(myChatbot, err, false)
			os.Exit(0)
		}
	}
}

func checkError(chatbot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}

	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
