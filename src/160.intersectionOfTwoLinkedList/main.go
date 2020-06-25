package main

import (
	"fmt"
)

func main() {
	a := &ListNode{
		Val:  0,
		Next: nil,
	}
	a.Next = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	b := &ListNode{
		Val:  3,
		Next: a.Next.Next.Next,
	}
	node := getIntersectionNode(a, b)
	fmt.Println("intersection node:", node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	walkerA := headA
	walkerB := headB
	reachedA := false
	reachedB := false
	for {
		if walkerA == walkerB {
			return walkerA
		}
		if walkerA.Next == nil {
			if reachedA {
				// linked list A reach list B's tail
				return nil
			}
			// last node of linked list A
			walkerA = headB
			reachedA = true
		} else {
			walkerA = walkerA.Next
		}
		if walkerB.Next == nil {
			if reachedB {
				return nil
			}
			walkerB = headA
			reachedB = true
		} else {
			walkerB = walkerB.Next
		}
	}
}
