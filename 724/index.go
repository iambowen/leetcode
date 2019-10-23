package main

import "fmt"

func half(a int) int {
	if a%2 == 0 || a > 0 {
		return a / 2
	}
	return a/2 - 1
}

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	a := half(sum)
	b := half(sum)
	i := 0
	for ; i < len(nums); i++ {
		if a == half(nums[i]) {
			break
		}
		a -= nums[i]
	}
	if i == len(nums) {
		return -1
	}
	b = b - half(nums[i])
	for j := i + 1; j < len(nums); j++ {
		b -= nums[j]
	}
	if b == 0 {
		return i
	}
	return -1
}
func main() {
	// a := []int{-1, -1, -1, -1, -1, 0}
	// a := []int{-1, -1, -1, 0, 1, 1} //0
	a := []int{-1, -1, -1, 1, 1, 1} //0
	// a := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(a))
}
