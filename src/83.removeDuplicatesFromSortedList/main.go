package main

import "fmt"

func main() {
	ln := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}
	ln = &ListNode{}
	r := deleteDuplicates(ln)
	for {
		if r == nil {
			break
		}
		fmt.Println(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for {
		if current.Next == nil {
			break
		}
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
