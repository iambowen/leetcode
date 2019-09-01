// https://leetcode-cn.com/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int64, target int64) []int64 {
	result := make([]int64, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				result = append(result, int64(i))
				result = append(result, int64(j))
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
	var result = []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		for _, value := range strings.Split(scanner.Text(), " ") {
			a, _ := strconv.ParseInt(value, 10, 32)
			result = append(result, a)
		}
	}
	fmt.Println(result)
	fmt.Println(twoSum(result, 6))
}
