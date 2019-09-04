package main

import "fmt"

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
	for i := 5; i <= n; {
		count5 += countOf5(i)
		i += 5
	}

	return count5
}

func main() {
	var n int

	fmt.Scanf("%d", &n)
	fmt.Println(trailingZeroes(n))
}
