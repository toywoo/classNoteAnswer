package first_problem

import (
	"fmt"
	"math"
)

func QuadraticFormula() { // main
	var a float64
	var b float64
	var c float64
	var d float64
	var answer1 float64
	var answer2 float64

	fmt.Println("---------이차 방정식 근의 공식 계산기----------")
	fmt.Println("이차 방정식의 계수를 방정식의 오른차순으로 입력하여주세요")
	fmt.Scanf("%f %f %f", &a, &b, &c)

	d = b*b - 4*a*c
	answer1 = (-b + math.Sqrt(d)) / 2 * a
	answer2 = (-b - math.Sqrt(d)) / 2 * a

	fmt.Printf("근은 %f %f 입니다.\n", answer1, answer2)
}
