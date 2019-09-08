package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	reverseInput(input)

}

func reverseInput(input int) {
	if input < 10 {
		fmt.Print(input)
		return
	} else {
		reverseInput(input % 10)
		reverseInput(input / 10)
	}
}
