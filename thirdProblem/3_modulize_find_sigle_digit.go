package thirdproblem

import (
	"fmt"
	"math/rand"
	"strings"
)

const ( // string은 위험하기 때문에 enum 선언을 하는 편이 좋기도 하고 이렇게 하면 알아보기 쉬움
	YES    = "y"
	NO     = "n"
	CHANCE = 3
)

func getName(name *string) {
	fmt.Print("What is your name \n>> ")
	fmt.Scanf("%s\n", name)

}

func isReady() {
	var isReady string

	for {
		fmt.Print("\nYou Ready?\nYes: Y, No: n\n>> ")
		fmt.Scanf("%s\n", &isReady)
		isReady = strings.ToLower(isReady)

		if isReady == YES {
			break
		} else if isReady == NO {
			return
		} else {
			fmt.Print("You input wrong character.\n\n")
		}
	}
}

func mainGame(name string) {
	var randNum int = rand.Intn(10)
	var inputNum int
	var usedChanceNum int = 0

	for usedChanceNum < CHANCE {
		switch usedChanceNum {
		case 0:
			fmt.Printf("%s, Choice single digit number!\n>> ", name)
		case 1:
			fmt.Printf("Fail! %s, Choice single digit number!\n>> ", name)
		case 2:
			fmt.Printf("Fail! %s, Choice single digit number! Last Chance.\n>> ", name)
		}

		_, err := fmt.Scanf("%d\n", &inputNum)

		if err != nil {
			fmt.Println("You input wrong character.")
			return
		}

		if inputNum == randNum {
			fmt.Printf("%s, Congratulation! You found hidden number.\n", name)
			return
		}

		usedChanceNum++
	}

	fmt.Printf("%s, Fail! That's too bad\n", name)
}

func FindSingleDigit() {
	var name string

	getName(&name)

	isReady()

	mainGame(name)

}
