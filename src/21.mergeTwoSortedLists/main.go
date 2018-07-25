package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toArray() []int {
	list := make([]int, 0, 10)
	list = append(list, l.Val)
	if l.Next != nil {
		list = append(list, l.Next.toArray()...)
	}
	return list
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 4}}}}
	l2 := &ListNode{2, &ListNode{2, &ListNode{4, &ListNode{Val: 6}}}}
	fmt.Println(l1.toArray(), l2.toArray())
	r := mergeTwoLists(l1, l2)
	fmt.Println(r.toArray())
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	p := r
	for {
		tmp := &ListNode{}
		if l1 != nil {
			if l2 != nil {
				if l1.Val <= l2.Val {
					tmp.Val = l1.Val
					l1 = l1.Next
				} else {
					tmp.Val = l2.Val
					l2 = l2.Next
				}
			} else {
				tmp.Val = l1.Val
				l1 = l1.Next
			}
			p.Next = tmp
			p = p.Next
		} else {
			if l2 != nil {
				tmp.Val = l2.Val
				l2 = l2.Next
				p.Next = tmp
				p = p.Next
			} else {
				break
			}
		}
	}
	return r.Next
}
