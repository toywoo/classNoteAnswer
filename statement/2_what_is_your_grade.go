package statement

import "fmt"

func WhatIsYourGrade() { // main
	var score uint

	fmt.Scanf("%d", &score)

	score /= 10

	switch score {
	case 10:
		fmt.Print("A")
	case 9:
		fmt.Print("A")
	case 8:
		fmt.Print("B")
	case 7:
		fmt.Print("C")
	case 6:
		fmt.Print("D")
	default:
		fmt.Print("F")
	}

	fmt.Print("\n") // main.go 때문에 존재하는 것
}
