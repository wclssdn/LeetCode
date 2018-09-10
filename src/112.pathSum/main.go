package main

import "fmt"

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}, &TreeNode{2, &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}, &TreeNode{Val: 3}}}
	r := hasPathSum(tree, 7)
	fmt.Println("has path sum:", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
