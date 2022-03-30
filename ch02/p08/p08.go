package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, l1}
	l3 := &ListNode{3, l2}
	l4 := &ListNode{4, l3}
	l5 := &ListNode{5, l4}
	l6 := &ListNode{6, l5}
	l7 := &ListNode{7, l6}
	l8 := &ListNode{8, l7}
	l9 := &ListNode{9, l8}
	l10 := &ListNode{10, l9}
	l11 := &ListNode{11, l10}
	l1.Next = l8
	fmt.Println(findLoopStart(l11))
}

func findLoopStart(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}

	fast, slow := l, l

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = l

	for l != nil && slow != nil {
		l = l.Next
		slow = slow.Next

		if l == slow {
			return l
		}
	}

	return nil
}
