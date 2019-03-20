package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("release-logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	if err != nil {
		Error.Println("日志打开失败")
	}

	debuggerLog := log.New(logFile, "Info:", log.LstdFlags)

	debuggerLog.Println("Hello World")
}
