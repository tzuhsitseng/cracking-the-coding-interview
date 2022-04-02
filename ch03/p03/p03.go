package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeStack struct {
	stack []*ListNode
}

type ListNodeStacks struct {
	stacks []*ListNodeStack
}

func NewListNodeStack() *ListNodeStack {
	return &ListNodeStack{
		stack: make([]*ListNode, 0),
	}
}

func NewListNodeStacks() *ListNodeStacks {
	return &ListNodeStacks{
		stacks: make([]*ListNodeStack, 0),
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

func (s *ListNodeStack) isFull() bool {
	return len(s.stack) == 4
}

func (s *ListNodeStack) peek() *ListNode {
	if s.isEmpty() {
		return nil
	}

	return s.stack[len(s.stack)-1]
}

func (s *ListNodeStacks) push(element *ListNode) {
	if len(s.stacks) == 0 {
		ns := NewListNodeStack()
		s.stacks = []*ListNodeStack{ns}
	}

	stack := s.stacks[len(s.stacks)-1]
	if len(s.stacks) > 0 && !stack.isFull() {
		stack.push(element)
	} else {
		ns := NewListNodeStack()
		ns.push(element)
		s.stacks = append(s.stacks, ns)
	}
}

func (s *ListNodeStacks) pop() *ListNode {
	if len(s.stacks) == 0 {
		return nil
	}

	stack := s.stacks[len(s.stacks)-1]
	p := stack.pop()
	if stack.isEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return p
}

func (s *ListNodeStacks) popAt(index int) *ListNode {
	if len(s.stacks) == 0 {
		return nil
	}
	return s.leftShift(index)
}

func (s *ListNodeStacks) leftShift(index int) *ListNode {
	stack := s.stacks[index]
	result := stack.pop()

	if stack.isEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1]
	} else {
		p := s.leftShift(index + 1)
		stack.push(p)
	}

	return result
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	stacks := NewListNodeStacks()
	stacks.push(n1)
	stacks.push(n2)
	stacks.push(n3)
	stacks.push(n4)
	stacks.push(n5)
	fmt.Println(stacks)

	stacks.popAt(0)
	fmt.Println(stacks)

	stacks.pop()
	stacks.pop()
	fmt.Println(stacks)
}
