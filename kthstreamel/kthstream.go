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
	return pq[i] < pq[j]
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

type KthLargest struct {
	k          int
	pq         PriorityQueue
	KthLargest int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{0, PriorityQueue{}, 0}
	kth.k = k
	for i := 0; i < len(nums); i++ {
		heap.Push(&kth.pq, nums[i])
	}
	for i := 0; kth.pq.Len() > k; i++ {
		heap.Pop(&kth.pq)
	}
	fmt.Println(kth.pq)
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.pq.Len() < this.k {
		heap.Push(&(this.pq), val)
		return this.pq[0]
	}
	if val <= this.pq[0] {
		return this.pq[0]
	}
	heap.Pop(&this.pq)
	heap.Push(&(this.pq), val)
	return this.pq[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)
	fmt.Println(kthLargest.pq)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
	return
}
