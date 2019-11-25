package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	digits := "+-*/"
	return !strings.Contains(digits, s)
}

func evalRPN(s []string) int {
	stack := []int{}

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num, _ := strconv.Atoi(s[i])
			stack = append(stack, num)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			result := calculate(a, b, s[i])
			stack = append(stack, result)
		}
	}
	return stack[0]
}

func calculate(a, b int, calculator string) int {
	switch calculator {
	case "+":
		return a + b
	case "*":
		return a * b
	case "/":
		return a / b
	case "-":
		return a - b
	default:
		return -1
	}
}

func main() {
	// s := []string{"2", "1", "+", "3", "*"}
	a := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// fmt.Println(evalRPN(s))
	fmt.Println(evalRPN(a))
}
