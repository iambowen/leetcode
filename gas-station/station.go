package main

type Station struct {
	d2Dest int
	gas    int
}

type PriorityQueue []Station

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].d2Dest-pq[i].gas > pq[j].d2Dest-pq[j].gas
}

func (pq *PriorityQueue) Push(x interface{}) {
	station := x.(Station)
	*pq = append(*pq, station)
}
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	station := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return station
}
func main() {
	stations := []Station{
		{15, 10},
		{11, 5},
		{5, 2},
		{4, 4},
	}
	distance, currentGas := 25, 10

	return
}
