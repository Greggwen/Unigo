package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var ret []byte
	var cmd *exec.Cmd
	var err error

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	input = strings.Trim(string(input), "\n")

	switch input {
	case "jinmao":
		cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao && git pull'`)
	case "jinmao-management":
		cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-management && git pull'`)
	case "jinmao-static":
		cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-static && git pull'`)
	case "jinmao-statistics":
		cmd = exec.Command("/bin/bash", "-c", `ansible webserver -m shell -a 'cd /home/mosh/webroot/jinmao-statistics && git pull'`)
	default:
		os.Exit(1)
	}

	if ret, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(ret))

}
