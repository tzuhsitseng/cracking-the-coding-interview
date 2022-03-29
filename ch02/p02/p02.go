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
	//fmt.Println(returnKthToLastByRec(l5, 5))
	fmt.Println(returnKthToLastByIte(l5, 2))
}

func returnKthToLastByRec(l *ListNode, k int) *ListNode {
	_, n := findKth(l, k)
	return n
}

func returnKthToLastByIte(l *ListNode, k int) *ListNode {
	p1, p2 := l, l

	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
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
