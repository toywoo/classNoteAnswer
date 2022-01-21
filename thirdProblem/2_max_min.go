package thirdproblem

import (
	"fmt"
)

func findMax(strNums []int) {
	var maxValue int = strNums[0]

	for i := 0; i < len(strNums); i++ {
		if maxValue < strNums[i] {
			maxValue = strNums[i]
		}
	}

	fmt.Println("\n최댓값: ", maxValue)
}

func findMin(strNums []int) {
	var minValue int = strNums[0]

	for i := 0; i < len(strNums); i++ {
		if minValue > strNums[i] {
			minValue = strNums[i]
		}
	}

	fmt.Println("\n", "최솟값: ", minValue)
}

func FindMaxMin() { // main
	var isMaxMin string
	var inpLeng int
	var inpNum int
	var inputSli []int
	fmt.Scanf("%s %d\n", &isMaxMin, &inpLeng)

	for i := 0; i < inpLeng; i++ {
		fmt.Scanf("%d\n", &inpNum)
		inputSli = append(inputSli, inpNum)
	}

	if isMaxMin == "max" {
		findMax(inputSli)
	} else {
		findMin(inputSli)
	}
}
