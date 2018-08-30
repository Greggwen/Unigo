package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var firstInput1, secondInput1, sum float64
var operator string

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input the first number:")
	firstInput, err := inputReader.ReadString('\n')
	//firstInput, err := inputReader.ReadRune()
	//fmt.Println(strconv.ParseFloat(firstInput, 64))
	if err != nil {
		os.Exit(1)
	}
	firstInput = strings.Replace(firstInput, "\n", "", -1)

	// 得到第一个参数，并转换成float64类型
	firstInput1, err := strconv.ParseFloat(firstInput, 64)

	//var supportOperator  = [4]string{"+", "-", "*", "/"}

	for {
		// 获取操作符【+-*/】
		fmt.Println("Please select the operator of [+,-,*,/]:")
		operator, err := inputReader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		operator = strings.Replace(operator, "\n", "", -1)

		switch operator {
		case "+":
			break
		case "-":
			break
			//os.Exit(1)
		case "*":
			break
			//os.Exit(1)
		case "/":
			break
			//os.Exit(1)
		default:
			fmt.Println("invalid operator")
			continue
		}

		fmt.Println("Please input the second number:")
		secondInput, err := inputReader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		secondInput = strings.Replace(secondInput, "\n", "", -1)
		// 得到第二个参数，并转换成float64类型
		secondInput1, err = strconv.ParseFloat(secondInput, 64)

		if operator == "+" {
			sum = firstInput1 + secondInput1
		} else if operator == "-" {
			sum = firstInput1 - secondInput1
		} else if operator == "*" {
			sum = firstInput1 * secondInput1
		} else if operator == "/" {
			sum = firstInput1 / secondInput1
		}

		fmt.Printf("%.2f %s %.2f = %.2f", firstInput1, operator, secondInput1, sum)
		os.Exit(0)
	}
}
