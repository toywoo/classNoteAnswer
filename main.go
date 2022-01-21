package main

import (
	first_problem "classNoteAnswer/firstProblem"
	"classNoteAnswer/statement"
	thirdproblem "classNoteAnswer/thirdProblem"
	"fmt"
)

// func onStatementFuncs(fileNum int) {

// }

// func onFunction(foldNum int64, fileNum int) {

// }

func main() {
	fmt.Print(`

예시 답안 코드 실행기 

사용방법
	폴더_번호 파일_번호 이런 방식으로 입력하시면 작동됩니다.
		  폴더 번호		| 파일 번호
		1 firstProblem	 8_qudratic_formula.go

		2 statement		 1_odd_or_even.go
						 2_what_is_your_grade.go
						 3_printing_star.go
						 4_mini_console_program.go
						 5_find_single_digit.go
		
		3 thirdProblem	 2_max_min.go
						 3_modulize_find_sigle_digit.go
						 6_grow_dog.go
						 7_url_decomposer.go
						 8_linked_list.go
	예시 | 2 statement > 1_odd_or_even.go 실행시
	>> exec 2 1
		
	종료를 원할시 exit or 강제종료(ctrl+C) 문제 발생시 ReadMe를 참고 해주세요.
`)

	var comm string
	var fileNum int
	var foldNum int64

	for {
		fmt.Print(">> ")
		fmt.Scanf("%s %d %d\n", &comm, &foldNum, &fileNum)

		if comm == "exit" {
			return
		} else if comm == "exec" {
			switch foldNum {
			case 1:
				first_problem.QuadraticFormula()
			case 2:
				switch fileNum {
				case 1:
					statement.OddOrEven()
				case 2:
					statement.WhatIsYourGrade()
				case 3:
					statement.PrintingStar()
				case 4:
					statement.MiniConsoleProgram()
				case 5:
					statement.FindSingleDigit()
				default:
					fmt.Println("명령어를 잘못 입력하셨습니다.")
				}
			case 3:
				switch fileNum {
				case 2:
					thirdproblem.FindMaxMin()
				case 3:
					thirdproblem.FindSingleDigit()
				case 6:
					thirdproblem.GrowDog()
				case 7:
					thirdproblem.URLDecomposer()
				case 8:
					thirdproblem.LinkedList()
				default:
					fmt.Println("명령어를 잘못 입력하셨습니다.")
				}
			default:
				fmt.Println("명령어를 잘못 입력하셨습니다.")
			}
			comm = ""
		} else if comm == "" {
			fmt.Println("입력을 받는 프로그램 작동시 발생하는 버퍼 오류 제거")
		} else {
			fmt.Println("명령어를 잘못 입력하셨습니다.")
		}
	}
}
