package main

import (
	"fmt"
)

func main() {
	// var x list.List
	// x.PushBack(1)
	// x.PushBack(2)
	// x.PushBack(3)
	// for e := x.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value.(int))
	// }
	// var a = []int{1, 2, 3}
	// a, value := a[0:len(a)-1], a[len(a)-1]
	// fmt.Println(value)
	// fmt.Println(a)

	for i, ch := range []byte("abc") {
		fmt.Println(i, ch)
	}
}

// https://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go
