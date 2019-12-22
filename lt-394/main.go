package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	nextNum int
	thisStr string
}

func decodeString(s string) string {
	num := 0
	curStr := ""
	stack := []Pair{}

	for _, c := range s {
		if unicode.IsDigit(c) { // end of string
			n, _ := strconv.Atoi(string(c))
			num = num*10 + n
		} else if c == '[' { // end of number
			stack = append(stack, Pair{num, curStr})
			num, curStr = 0, ""
		} else if c == ']' {
			pair := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curStr = pair.thisStr + generate(pair.nextNum, curStr)
		} else {
			curStr = curStr + string(c)
		}
	}

	return curStr
}

func generate(n int, str string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += str
	}
	return result
}
func isNum(s string) bool {
	nums := "0123456789"
	return strings.Contains(nums, s)
}
func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}
