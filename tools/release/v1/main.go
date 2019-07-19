package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"log"
)

func main() {
	fmt.Println(os.Args[0])

	var ret []byte
	var cmd *exec.Cmd
	var err error

	//logFile, err := os.OpenFile("/home/mosh/console-log/release-logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logFile, err := os.OpenFile("/Users/simple.xull/DevOps/Code/local/golang/src/Unigo/tools/release/v1/release-logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close();

	if err != nil {
		log.Println("日志打开失败")
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your project name[jinmao, jinmao-management, jinmao-static, jinmao-statistics, jinmao-verification]:")
	input, err := inputReader.ReadString('\n')

	debuggerLog := log.New(logFile, "Info:", log.LstdFlags)

	debuggerLog.Println("===================================  START RELEASE ==========================");
	if err != nil {
		debuggerLog.Println("release fails!");
		debuggerLog.Println("===================================  END RELEASE ==========================");
		os.Exit(1)
	}

	input = strings.Trim(string(input), "\n")

	switch input {
	case "jinmao":
		fmt.Println("jinmao")
		cmd = exec.Command("/bin/bash", "-c", `ls -al`)
		//cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao && git pull'`)
	case "jinmao-management":
		fmt.Println("jinmao-management")
		cmd = exec.Command("/bin/bash", "-c", `ls -al`)
		//cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-management && git pull'`)
	case "jinmao-static":
		fmt.Println("jinmao-static")
		cmd = exec.Command("/bin/bash", "-c", `ls -al`)
		//cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-static && git pull'`)
	case "jinmao-statistics":
		fmt.Println("jinmao-statistics")
		cmd = exec.Command("/bin/bash", "-c", `ls -al`)
		//cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-statistics && git pull'`)
	case "jinmao-verification":
		cmd = exec.Command("/bin/bash", "-c", `ls -al`)
		//cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-verification && git pull'`)
	default:
		fmt.Println("The Project named " + input + " not exists!");
		debuggerLog.Println("The Project named " + input + " not exists!");
		debuggerLog.Println("===================================  END RELEASE ==========================");
		os.Exit(1)
	}

	if ret, err = cmd.Output(); err != nil {
		fmt.Println(err)
		debuggerLog.Println(err);
		debuggerLog.Println("===================================  END RELEASE ==========================");
		os.Exit(1)
	}

	debuggerLog.Println(string(ret));
	debuggerLog.Println("===================================  END RELEASE ==========================");
	fmt.Println(string(ret))

}
