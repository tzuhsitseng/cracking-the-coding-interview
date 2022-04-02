package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeStack struct {
	stack []*ListNode
}

type ListNodeStackWithMin struct {
	ListNodeStack
	minStack ListNodeStack
}

func NewListNodeStackWithMin() *ListNodeStackWithMin {
	return &ListNodeStackWithMin{
		ListNodeStack: ListNodeStack{},
		minStack:      ListNodeStack{},
	}
}

func (s *ListNodeStack) push(element *ListNode) {
	if s.stack == nil {
		s.stack = make([]*ListNode, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *ListNodeStack) pop() *ListNode {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return result
}

func (s *ListNodeStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *ListNodeStack) peek() *ListNode {
	if s.isEmpty() {
		return nil
	}

	return s.stack[len(s.stack)-1]
}

func (s *ListNodeStackWithMin) push(element *ListNode) {
	if s.stack == nil {
		s.stack = make([]*ListNode, 0)
	}
	s.stack = append(s.stack, element)

	min := s.min()
	if min == nil || element.Val <= s.min().Val {
		s.minStack.push(element)
	}
}

func (s *ListNodeStackWithMin) pop() *ListNode {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]

	min := s.min()
	if min != nil && result.Val == min.Val {
		s.minStack.pop()
	}

	return result
}

func (s *ListNodeStackWithMin) min() *ListNode {
	if s.minStack.isEmpty() {
		return nil
	}

	return s.minStack.peek()
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}

	s := NewListNodeStackWithMin()
	s.push(n3)
	s.push(n2)
	s.push(n1)
	fmt.Println(s.min())
	s.pop()
	fmt.Println(s.min())
	s.push(n1)
	fmt.Println(s.min())
}
