package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	sums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sums[i] = i + nums[i]
	}
	return canJump2(0, sums)
}

func canJump2(start int, sums []int) bool {
	if sums[start] >= len(sums)-1 {
		return true
	}
	max := sums[start]
	index := start
	for j := start + 1; j <= sums[start]; j++ {
		if max <= sums[j] {
			max = sums[j]
			index = j
		}
	}
	if index == start {
		return false
	}
	return canJump2(index, sums)
}

func main() {
	// sums := []int{2, 3, 1, 1, 4} //true
	// nums := []int{3, 2, 1, 0, 4} //false
	// nums := []int{0} //true
	// nums := []int{1, 2, 3} //true
	nums := []int{1, 1, 1, 0} //true
	fmt.Println(canJump(nums))
}
