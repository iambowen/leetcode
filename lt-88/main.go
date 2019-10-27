package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	right := m + n - 1
	n1 := m - 1
	n2 := n - 1
	for n1 >= 0 && n2 >= 0 {
		if nums1[n1] <= nums2[n2] {
			nums1[right], nums2[n2] = nums2[n2], nums1[right]
			right--
			n2--
		} else if nums1[n1] > nums2[n2] {
			nums1[n1], nums2[n2] = nums2[n2], nums1[n1]
			nums1[right], nums2[n2] = nums2[n2], nums1[right]
			right--
			n1--
		}
	}
	for i := n1; i > 0; i-- {
		if nums1[i] < nums1[i-1] {
			nums1[i], nums1[i-1] = nums1[i-1], nums1[i]
		}
	}

	// fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m, n := 3, 3
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)
}
