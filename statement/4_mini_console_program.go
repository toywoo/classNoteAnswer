package statement

import (
	"fmt"
	"strconv"
	"strings"
)

func MiniConsoleProgram() { // main
	var comm string
	var op1 string
	var op2 string
	var op3 string
	var numOp2 int64
	var numOp3 int64
	var convertError1 error
	var convertError2 error

	for {
		fmt.Print(">> ")
		fmt.Scanf("%s %s %s %s\n", &comm, &op1, &op2, &op3)

		comm = strings.ToLower(comm)

		if comm == "calc" {
			op1 = strings.ToLower(op1)
			numOp2, convertError1 = strconv.ParseInt(op2, 10, 64)
			numOp3, convertError2 = strconv.ParseInt(op3, 10, 64)

			if convertError1 != nil || convertError2 != nil {
				fmt.Println("숫자를 입력하세요.")
			} else {
				switch op1 {
				case "add":
					fmt.Println(numOp2 + numOp3)
				case "minu":
					fmt.Println(numOp2 - numOp3)
				case "time":
					fmt.Println(numOp2 * numOp3)
				case "mult":
					fmt.Println(numOp2 / numOp3)
				case "remi":
					fmt.Println(numOp2 % numOp3)
				default:
					fmt.Println("명령어를 다시 입력하세요.*")
				}
			}
		} else if comm == "print" {
			fmt.Println(op1, op2, op3)
		} else if comm == "exit" {
			fmt.Println("프로그램을 종료합니다.")
			return
		} else {
			fmt.Println("명령어를 다시 입력하세요.")
		}
	}
}
