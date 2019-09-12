package main

// https://leetcode-cn.com/problems/bulb-switcher/

import "fmt"

func bulbSwitch(n int) int {
	var count int
	for i := 1; i*i <= n; i++ {
		count++
	}
	return count
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n < 1 {
		panic("input error")
	}

	fmt.Println(bulbSwitch(n))
}
