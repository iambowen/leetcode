package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	fmt.Println(minStack.stack)
	minStack.Push(0)
	fmt.Println(minStack.stack)
	minStack.Push(-3)
	fmt.Println(minStack.stack)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	stack   []int
	pivotal int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.pivotal++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.pivotal-1]
	this.pivotal--
}

func (this *MinStack) Top() int {
	return this.stack[this.pivotal-1]
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for i := 1; i < this.pivotal; i++ {
		if min > this.stack[i] {
			min = this.stack[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
