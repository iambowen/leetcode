package main

import "fmt"

// return value of n!, 0<n<10000

func main() {
	results := [10000]int{0}
	results[0] = 1
	var n int
	fmt.Scanf("%d", &n)

	for i := 2; i <= n; i++ {
		temp := 0
		for j := 0; j < 10000; j++ {
			a := results[j]*i + temp
			results[j] = a % 10
			temp = a / 10
		}
	}

	nonZeroIndex := 0
	for j := 9999; j >= 0; j-- {
		if results[j] != 0 {
			nonZeroIndex = j
			break
		}
	}

	for i := nonZeroIndex; i >= 0; i-- {
		fmt.Print(results[i])
	}
	fmt.Println("")
}
