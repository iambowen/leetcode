package main

import "fmt"

func dailyTemperatures(T []int) []int {
	results := make([]int, len(T))
	results[len(results)-1] = 0
	stack := []int{}
	stack = append(stack, T[len(T)-1])

	for i := len(T) - 2; i >= 0; i-- {
		count := 0
		yes := false
		for j := len(stack) - 1; j >= 0; j-- {
			count++
			if stack[j] > T[i] {
				yes = true
				break
			}
		}
		stack = append(stack, T[i])
		if !yes {
			results[i] = 0
		} else {
			results[i] = count
		}
	}
	return results
}
func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(t))
}
