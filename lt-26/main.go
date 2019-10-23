package main

//双指针法

// func removeDuplicates(nums []int) int {
// 	actualLast := len(nums) - 1
// 	for pivotal := 0; pivotal <= actualLast; pivotal++ {
// 		for i := pivotal + 1; i <= actualLast; {
// 			if nums[pivotal] == nums[i] {
// 				temp := nums[i]
// 				nums = append(nums[:i], nums[i+1:]...)
// 				nums = append(nums, temp)
// 				actualLast--
// 			} else {
// 				i++
// 			}
// 		}
// 	}
// 	return actualLast + 1
// }

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(a)
}
