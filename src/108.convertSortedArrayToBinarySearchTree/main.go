package main

import "fmt"

func main() {
	array := []int{-10, -3, 0, 5, 9}
	array = []int{1, 3}
	array = []int{1, 2, 3, 4, 5}
	r := sortedArrayToBST(array)
	printTree(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(tree *TreeNode) {
	if tree == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("%d ", tree.Val)
	printTree(tree.Left)
	printTree(tree.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	tree := &TreeNode{}
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		tree.Val = nums[0]
		return tree
	}
	m := l / 2
	tree.Val = nums[m]

	if m > 0 {
		tree.Left = &TreeNode{}
		insert(tree.Left, nums[:m])
	}
	if m+1 < l {
		tree.Right = &TreeNode{}
		insert(tree.Right, nums[m+1:])
	}

	return tree
}

func insert(tree *TreeNode, nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}
	if l == 1 {
		tree.Val = nums[0]
		return
	}
	m := l / 2
	tree.Val = nums[m]

	if m > 0 {
		tree.Left = &TreeNode{}
		insert(tree.Left, nums[:m])
	}
	if m+1 < l {
		tree.Right = &TreeNode{}
		insert(tree.Right, nums[m+1:])
	}
}
