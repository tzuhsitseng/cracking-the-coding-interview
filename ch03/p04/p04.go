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

type ListNodeStackQueue struct {
	push, pop *ListNodeStack
}

func NewListNodeStackQueue() *ListNodeStackQueue {
	return &ListNodeStackQueue{
		push: NewListNodeStack(),
		pop:  NewListNodeStack(),
	}
}

func (q *ListNodeStackQueue) add(element *ListNode) {
	if q.push.stack == nil {
		q.push.stack = make([]*ListNode, 0)
	}
	q.push.push(element)
}

func (q *ListNodeStackQueue) remove() *ListNode {
	if !q.pop.isEmpty() {
		return q.pop.pop()
	}

	for !q.push.isEmpty() {
		p := q.push.pop()
		q.pop.push(p)
	}
	return q.pop.pop()
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	q := NewListNodeStackQueue()
	q.add(n1)
	q.add(n2)
	q.add(n3)
	q.add(n4)
	q.add(n5)
	fmt.Println(q)

	q.remove()
	q.remove()
	fmt.Println(q)
}
