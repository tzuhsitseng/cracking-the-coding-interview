package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(l *ListNode) *ListNode {
	exists := map[int]bool{}
	previous := &ListNode{}
	iter := l

	for iter != nil {
		if _, ok := exists[iter.Val]; ok {
			previous.Next = iter.Next
		} else {
			exists[iter.Val] = true
			previous = iter
		}
		iter = iter.Next
	}

	return l
}

func deleteDuplicatesWithoutHashMap(l *ListNode) *ListNode {
	cur := l

	for cur != nil {
		runner := cur

		for runner.Next != nil {
			if runner.Next.Val == cur.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		cur = cur.Next
	}

	return l
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, l1}
	l3 := &ListNode{3, l2}
	l4 := &ListNode{2, l3}
	l5 := &ListNode{4, l4}
	result := deleteDuplicates(l5)
	fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val)
	result = deleteDuplicatesWithoutHashMap(l5)
	fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val)
}
