package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 1, 2
	for ; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[0 : i+1])
	return i + 1

}
func main() {
	//         0, 0, 1, 1, 2, 2, 3, 3, 4, 4
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(a))
	fmt.Println(removeDuplicates(b))

}
