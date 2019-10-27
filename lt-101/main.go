package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(right *TreeNode, left *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}
	if right == nil || left == nil || right.Val != left.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// func isSymmetricBFS(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	temp := []*TreeNode{}

// 	temp = append(temp, root.Left, root.Right)
// 	for len(temp) != 0 {
// 		fmt.Println(len(temp))
// 		if len(temp)%2 != 0 {
// 			return false
// 		}
// 		a := temp
// 		temp = []*TreeNode{}
// 		i := 0
// 		j := len(a) - 1
// 		for i < j {
// 			if (a[i] == nil && a[j] != nil) || (a[i] != nil && a[j] == nil) || a[i].Val != a[j].Val {
// 				return false
// 			}
// 			i++
// 			j--
// 		}
// 		for k := 0; k < len(a); k++ {
// 			if a[k].Right != nil || a[k].Left != nil {
// 				temp = append(temp, a[k].Left, a[k].Right)
// 			}
// 		}
// 	}
// 	return true
// }

func main() {
	// 1,2,2,3,4,4,3
	// 1,2,2,null,3,null,3
	root := TreeNode{1, nil, nil}
	node1 := TreeNode{2, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node4 := TreeNode{3, nil, nil}
	root.Left = &node1
	root.Right = &node2
	node1.Left = nil
	node1.Right = &node3
	node2.Left = nil
	node2.Right = &node4
	fmt.Println(isSymmetricIter(&root))
}
