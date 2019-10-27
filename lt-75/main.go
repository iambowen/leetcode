package main

import "fmt"

func sortColors(nums []int) {
	right := len(nums) - 1
	left := 0
	current := 0
	for current < right {
		if nums[current] == 0 {
			nums[current], nums[left] = nums[left], nums[current]
			left++
			current++
		} else if nums[current] == 2 {
			nums[current], nums[right] = nums[right], nums[current]
			current++
			right--
		} else {
			current++
		}
	}
	// fmt.Println(nums)
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
}
