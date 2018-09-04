package main

import (
	"fmt"
	"math"
)

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, &TreeNode{Val: 4, Left: &TreeNode{Val: 33, Left: &TreeNode{Val: 33}}}}, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 3}}}
	r := isBalanced(tree)

	fmt.Println("balanced:", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n)
//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return true
//	}
//
//	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
//}
//
//func height(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	return 1 + int(math.Max(float64((height(root.Left))), float64(height(root.Right))))
//}

// dfs 方式O(n)
func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	if left == -1 {
		return -1
	}
	right := height(node.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return 1 + int(math.Max(float64(left), float64(right)))
}
