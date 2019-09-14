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
func lastStoneWeight(stones []int) int {
	pq := PriorityQueue{}
	for i := 0; i < len(stones); i++ {
		heap.Push(&pq, stones[i])
	}

	fmt.Println(pq)
	for i := 0; pq.Len() > 1; i++ {
		y, x := heap.Pop(&pq), heap.Pop(&pq)
		fmt.Println(pq)
		fmt.Println("y: ", y, " x: ", x)
		if x == y {
			continue
		} else if x.(int) < y.(int) {
			heap.Push(&pq, y.(int)-x.(int))
		} else {
			heap.Push(&pq, x.(int)-y.(int))
		}
	}
	if pq.Len() == 1 {
		return pq.Pop().(int)
	}
	return 0
}

func main() {
	// stones := []int{2, 7, 4, 1, 8, 1}
	stones := []int{4, 3, 4, 3, 2}
	// stones := []int{3, 7, 2}
	fmt.Println(stones)
	fmt.Println(lastStoneWeight(stones))
}
