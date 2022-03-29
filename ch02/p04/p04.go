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
	l4 := &ListNode{2, l3}
	l5 := &ListNode{4, l4}
	result := partition(l5, 3)
	fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val, result.Next.Next.Next.Next.Val)
}

func partition(l *ListNode, x int) *ListNode {
	head, tail := l, l

	for l != nil {
		next := l.Next

		if l.Val < x {
			l.Next = head
			head = l
		} else {
			tail.Next = l
			tail = l
		}

		l = next
	}

	tail.Next = nil
	return head
}
