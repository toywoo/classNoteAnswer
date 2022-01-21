package thirdproblem

import (
	"fmt"
	"strings"
)

func URLDecomposer() { // main
	var url string

	fmt.Scanf("%s\n", &url)

	url = strings.ReplaceAll(url, "https://", "")
	url = strings.ReplaceAll(url, "http://", "")

	for index, value := range strings.Split(url, "/") {
		fmt.Println(index, value)
	}
}
