package main

import "fmt"

func main() {
	q := &TreeNode{1, &TreeNode{2, nil, &TreeNode{Val: 4}}, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 3}}}
	r := isSymmetric(q)
	fmt.Println("r", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}

	return isSame(root.Left, root.Right)
}
func isSame(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if a.Left == nil && b.Right == nil && a.Val == b.Val {
		return isSame(a.Right, b.Left)
	}
	if a.Right == nil && b.Left == nil && a.Val == b.Val {
		return isSame(a.Left, b.Right)
	}
	return isSame(a.Right, b.Left) && isSame(a.Left, b.Right)
}
