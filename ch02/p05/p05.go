package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{8, nil}
	l2 := &ListNode{2, l1}
	l3 := &ListNode{9, l2}
	l4 := &ListNode{2, nil}
	l5 := &ListNode{4, l4}
	l6 := &ListNode{4, l5}
	result := addLists(l3, l6)
	fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val)
}

func addLists(l1, l2 *ListNode) *ListNode {
	var result *ListNode
	carry := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10

		if result == nil {
			result = &ListNode{
				Val: val,
			}
		} else {
			n := &ListNode{
				Val: val,
			}
			n.Next = result
			result = n
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		n := &ListNode{
			Val: carry,
		}
		n.Next = result
		result = n
	}

	return result
}
