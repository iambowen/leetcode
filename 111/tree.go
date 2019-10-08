package main

import "fmt"

func main() {
	node3 := TreeNode{3, nil, nil}
	node4 := TreeNode{4, nil, nil}
	node5 := TreeNode{5, nil, nil}
	node2 := TreeNode{2, &node4, &node5}
	node1 := TreeNode{1, &node2, &node3}

	fmt.Println(minDepth(&node1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	var nodes = []*TreeNode{}
	nodes = append(nodes, root)
	count := 0
	var bingo bool
	for len(nodes) != 0 && !bingo {
		tempNodes := nodes
		nodes = []*TreeNode{}
		count++
		for _, node := range tempNodes {
			if node.Left == nil && node.Right == nil {
				bingo = true
				break
			}
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}
	return count
}
