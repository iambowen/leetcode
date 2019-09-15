// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PriorityQueue) Push(x interface{}) {
	value := x.(int)
	*pq = append(*pq, value)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	value := (*pq)[n-1]
	fmt.Print("")
	*pq = (*pq)[:n-1]
	return value
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func findKthLargest(nums []int, k int) int {
	pq := PriorityQueue{}
	for i := 0; i < len(nums); i++ {
		heap.Push(&pq, nums[i])
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(&pq)
	}
	return heap.Pop(&pq).(int)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
