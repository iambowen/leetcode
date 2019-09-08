package main

import "fmt"

func trailingZeroes2(n int) int {
	if n < 5 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}

func countOf5(n int) int {
	var result int
	for n%5 == 0 {
		result++
		n = n / 5
	}
	return result
}

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	var count5 int
	for i := 5; i <= n; i = i + 5 {
		if i%25 == 0 {
			count5 += countOf5(i)
		} else {
			count5++
		}
	}

	return count5
}

func main() {
	var n int

	fmt.Scanf("%d", &n)
	// fmt.Println(trailingZeroes(n))
	fmt.Println(trailingZeroes2(n))
}
