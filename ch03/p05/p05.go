package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeStack struct {
	stack []*ListNode
}

func NewListNodeStack() *ListNodeStack {
	return &ListNodeStack{
		stack: make([]*ListNode, 0),
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

type SortedStack struct {
	data   *ListNodeStack
	buffer *ListNodeStack
}

func NewSortedStack() *SortedStack {
	return &SortedStack{
		data:   NewListNodeStack(),
		buffer: NewListNodeStack(),
	}
}

func (s *SortedStack) Sort() {
	for !s.data.isEmpty() {
		tmp := s.data.pop()

		for !s.buffer.isEmpty() && s.buffer.peek().Val > tmp.Val {
			s.data.push(s.buffer.pop())
		}

		s.buffer.push(tmp)
	}

	for !s.buffer.isEmpty() {
		s.data.push(s.buffer.pop())
	}
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	ss := NewSortedStack()
	ss.data.push(n4)
	ss.data.push(n2)
	ss.data.push(n3)
	ss.data.push(n5)
	ss.data.push(n1)
	fmt.Println(ss)
	ss.Sort()
	fmt.Println(ss)
}
