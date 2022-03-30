package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func push(stack []*ListNode, elements ...*ListNode) []*ListNode {
	return append(stack, elements...)
}

func pop(stack []*ListNode) (*ListNode, []*ListNode) {
	l := len(stack)
	if l > 0 {
		return stack[l-1], stack[:l-1]
	}
	return nil, nil
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, l1}
	l3 := &ListNode{3, l2}
	l4 := &ListNode{3, l3}
	l5 := &ListNode{2, l4}
	l6 := &ListNode{1, l5}
	fmt.Println(checkPalindrome(l6))
}

func checkPalindrome(l *ListNode) bool {
	stack := make([]*ListNode, 0)
	fast, slow := l, l

	for fast != nil && fast.Next != nil {
		stack = push(stack, slow)
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		var top *ListNode
		top, stack = pop(stack)
		if top.Val != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}
