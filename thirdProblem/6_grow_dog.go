package thirdproblem

import (
	"fmt"
	"math/rand"
	"time"
)

type Dog struct {
	hp      int
	clean   int
	satisfy int
	day     uint
	name    string
}

func adoptDog() *Dog {
	var name string

	fmt.Print("입양할 강아지의 이름을 입력 해주세요.\n>> ")
	fmt.Scanf("%s\n", &name)

	return &Dog{hp: 100, clean: 100, satisfy: 100, day: 0, name: name}
}

func (dog *Dog) reportToPolice(isReportPolice *bool) {
	if dog.hp+dog.clean+dog.satisfy < 151 {
		fmt.Println(dog.name, "가 당신을 매 저녁 해치는 사람으로 생각하여 경찰에 신고하였습니다. 당신은 이제 강아지를 키우지 못하게 됩니다.")
		*isReportPolice = true
	} else {
		dog.day++

		fmt.Println(dog.name, "에게 ", dog.day, "번째 날이 찾아왔습니다.")
		dog.currentState()
		*isReportPolice = false
	}
}

func (dog *Dog) eatFood() {
	if dog.hp < 91 && dog.satisfy < 91 {
		dog.hp += 10
		dog.satisfy += 10
	}
}

func (dog *Dog) washBody() {
	if dog.clean < 91 {
		dog.clean += 10
	}
}

func (dog *Dog) walkDog() {
	if dog.satisfy < 91 && dog.satisfy > 3 {
		dog.satisfy += 10
		dog.clean -= 3
	}
}

func (dog *Dog) currentState() {
	fmt.Println("<현재 상황>\t체력: ", dog.hp, "청결도: ", dog.clean, "만족도: ", dog.satisfy)
}

func (dog *Dog) randHate() {
	rand.Seed(time.Now().UnixNano())
	var randState int = rand.Int()
	rand.Seed(time.Now().UnixNano())
	var randAmount int = (rand.Int() % 5) * 10

	switch randState % 3 {
	case 0:
		dog.hp -= randAmount
	case 1:
		dog.clean -= randAmount
	case 2:
		dog.satisfy -= randAmount
	}
}

func GrowDog() { // main
	var isReportPolice bool = false
	var dog = adoptDog()

	var comm string

	for !isReportPolice {
		fmt.Print(">> ")
		fmt.Scanf("%s\n", &comm)

		if comm == "eat" {
			dog.eatFood()
		} else if comm == "wash" {
			dog.washBody()
		} else if comm == "walk" {
			dog.walkDog()
		} else if comm == "state" {
			dog.currentState()
		} else {
			fmt.Println("잘 못 입력 하셨습니다.")
		}

		dog.randHate()

		dog.reportToPolice(&isReportPolice)
	}

	fmt.Println("반려동물 키우기 종료됩니다.")
}
