package main

import (
	"fmt"
	"strconv"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	m := map[string]func(op int) int{}
	m["single"] = func(op int) int {
		return op
	}

	m["double"] = func(op int) int {
		return op * op
	}

	m["triple"] = func(op int) int {
		return op * op * op
	}

	fmt.Println("single : ", m["single"](2))
	fmt.Println("double : ", m["double"](2))
	fmt.Println("triple : ", m["triple"](2))

	s := "hello world"
	fmt.Println(s[1:2])

	for i := 0; i < len(s); i++ {
		fmt.Println("string:", s[i:i+1])
	}

	fmt.Println(len(strconv.Itoa(11111)))
}
