package statement

import "fmt"

func OddOrEven() { // main
	var numToFind uint

	fmt.Println("---------홀짝 판별기----------")
	fmt.Println(" 판별할 양의 정수를 입력하세요.")

	fmt.Scanf("%d", &numToFind)

	if (numToFind % 2) == 0 {
		fmt.Println("짝수")
	} else {
		fmt.Println("홀수")
	}
}
