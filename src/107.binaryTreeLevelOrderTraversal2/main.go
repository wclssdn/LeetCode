package main

import (
	"fmt"
)

func main() {
	tree := &TreeNode{1, &TreeNode{Val: 3}, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 5}}}
	r := levelOrderBottom(tree)
	fmt.Println("r:", r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ch := make(chan [2]int)
	done := make(chan struct{})
	over := make(chan struct{})
	max := 0
	result := [][]int{}
	result = append(result, []int{root.Val})
	go func() {
		for {
			select {
			case r := <-ch:
				if max < r[0] {
					max = r[0]
					for i := len(result); i <= max; i++ {
						result = append(result, []int{})
					}
				}
				result[r[0]] = append(result[r[0]], r[1])
			case <-done:
				close(over)
				return
			}

		}
	}()
	travsal(root, 1, ch)
	close(done)
	<-over
	tmp := [][]int{}
	for i := len(result) - 1; i >= 0; i-- {
		tmp = append(tmp, result[i])
	}
	return tmp
}

func travsal(root *TreeNode, level int, ch chan [2]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		ch <- [2]int{level, root.Left.Val}
		travsal(root.Left, level+1, ch)
	}
	if root.Right != nil {
		ch <- [2]int{level, root.Right.Val}
		travsal(root.Right, level+1, ch)
	}
}
