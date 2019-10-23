package main

import "fmt"

func match(left, right string) bool {
	return left == "(" && right == ")" || left == "{" && right == "}" || left == "[" && right == "]"
}

func isValid(s string) bool {
	if len(s)%2 != 0 && len(s) == 0 {
		return false
	}
	if s[:1] == "}" || s[:1] == ")" || s[:1] == "]" {
		return false
	}
	stack := make([]string, len(s))
	pivotal := 0
	for i := 0; i < len(s); {
		if pivotal == 0 {
			stack[pivotal] = s[i : i+1]
			pivotal++
			i++
		} else {
			switch match(stack[pivotal-1], s[i:i+1]) {
			case true:
				stack[pivotal-1] = ""
				i++
				pivotal--
			case false:
				stack[pivotal] = s[i : i+1]
				i++
				pivotal++
			}
		}

	}

	return pivotal == 0
}

func main() {
	// s := "()[]{}"
	// s := "()[]{}"
	// s := "([)]"
	s := "{[]}"
	fmt.Println(isValid(s))
}
