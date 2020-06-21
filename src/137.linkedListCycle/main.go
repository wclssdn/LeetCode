package main

import "fmt"

// https://leetcode.com/problems/linked-list-cycle/
func main() {
	l := &ListNode{
		Val:  3,
		Next: nil,
	}
	l.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  -4,
				Next: l,
			},
		},
	}

	cycle := hasCycle(l)
	fmt.Println("has cycle:", cycle)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowWalker := head
	fastWalker := head.Next
	for {
		if slowWalker == fastWalker {
			return true
		}
		if slowWalker == nil || fastWalker == nil || fastWalker.Next == nil {
			return false
		}
		slowWalker = slowWalker.Next
		fastWalker = fastWalker.Next.Next
	}
}
