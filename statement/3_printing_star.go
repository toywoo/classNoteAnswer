package statement

import "fmt"

func PrintingStar() { // main
	fmt.Print("1\n")

	for i := 1; i < 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	fmt.Print("\n2\n")

	for i := 6; 0 < i; i-- {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	fmt.Print("3\n")

	for i := 1; i < 6; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
