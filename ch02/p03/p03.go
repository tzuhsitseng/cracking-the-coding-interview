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
	_ = &ListNode{4, l4}
	fmt.Println(deleteNode(l3))
	fmt.Println(deleteNode(l1))
}

func deleteNode(del *ListNode) bool {
	if del == nil || del.Next == nil {
		return false
	}

	next := del.Next
	del.Val = next.Val
	del.Next = next.Next
	return true
}
