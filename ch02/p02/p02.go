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
	fmt.Println(returnKthToLast(l5, 5))
}

func returnKthToLast(l *ListNode, k int) *ListNode {
	_, n := findKth(l, k)
	return n
}

func findKth(l *ListNode, k int) (int, *ListNode) {
	if l == nil {
		return 0, nil
	}

	idx, node := findKth(l.Next, k)
	if idx == k {
		return k, node
	}
	return idx + 1, l
}
