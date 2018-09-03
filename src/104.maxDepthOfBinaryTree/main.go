package main

import (
	"fmt"
	"math"
)

func main() {
	tree := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	r := maxDepth(tree)

	fmt.Println("max depth:", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return maxDepth2(root, 0)
//}
//
//func maxDepth2(node *TreeNode, depth int) int {
//	if node == nil {
//		return depth
//	}
//	return int(math.Max(float64(maxDepth2(node.Left, depth+1)), float64(maxDepth2(node.Right, depth+1))))
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
