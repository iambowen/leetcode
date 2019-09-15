package main

//issue: https://leetcode-cn.com/problems/top-k-frequent-elements/
import (
	"container/heap"
	"fmt"
)

type Item struct {
	key       int
	frequence int //词频
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequence > pq[j].frequence
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	q := PriorityQueue{}
	// 入堆
	for key, frequence := range m {
		heap.Push(&q, &Item{key: key, frequence: frequence})
	}

	var result []int
	// pop k次
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}

func main() {
	nums := []int{2, 7, 4, 1, 8, 1}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
