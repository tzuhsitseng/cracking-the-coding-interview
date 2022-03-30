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
	l4 := &ListNode{3, l3}
	l5 := &ListNode{2, l2}
	l6 := &ListNode{1, l5}
	fmt.Println(findIntersection(l4, l6))
}

func findIntersection(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var l1Tail, l2Tail *ListNode
	var l1Len, l2Len int

	l1Tail, l1Len = getTailAndSize(l1)
	l2Tail, l2Len = getTailAndSize(l2)

	if l1Tail != l2Tail {
		return nil
	}

	if l1Len > l2Len {
		for i := 0; i < l1Len-l2Len; i++ {
			l1 = l1.Next
		}
	} else {
		for i := 0; i < l2Len-l1Len; i++ {
			l2 = l2.Next
		}
	}

	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return nil
}

func getTailAndSize(l *ListNode) (*ListNode, int) {
	var tail *ListNode
	var length int

	iter := l
	for iter != nil {
		length++
		if iter.Next == nil {
			tail = iter
		}
		iter = iter.Next
	}
	return tail, length
}
