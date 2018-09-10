package main

import (
	"fmt"
	"math"
)

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}, &TreeNode{2, &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}, &TreeNode{Val: 3}}}
	r := minDepth(tree)
	fmt.Println("min depth:", r)
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
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lh := dspHeight(root.Left)
	rh := dspHeight(root.Right)
	if lh == 0 {
		return 1 + int(rh)
	}
	if rh == 0 {
		return 1 + int(lh)
	}
	return 1 + int(math.Min(lh, rh))
}

func dspHeight(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	if node.Left == nil || node.Right == nil {
		return 1 + math.Max(dspHeight(node.Left), dspHeight(node.Right))
	}
	return 1 + math.Min(dspHeight(node.Left), dspHeight(node.Right))
}
