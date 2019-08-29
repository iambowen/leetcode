// https://leetcode-cn.com/problems/two-sum/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				result = append(result, i)
				result = append(result, j)
			}
		}

	}
	return result
}

func main() {
	// var result [20]byte
	// var n byte

	// _, err := os.Stdin.Read(result[:])
	// if err != nil {
	// 	fmt.Println("read error:", err)
	// 	return
	// }
	// _, err := os.Stdin.Read(n)
	// if err != nil {
	// 	fmt.Println("read error:", err)
	// 	return
	// }
	var result = []int{3, 2, 4}

	fmt.Println(twoSum(result, 6))
}
