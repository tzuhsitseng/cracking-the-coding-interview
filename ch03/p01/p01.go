package main

import (
	"errors"
	"fmt"
)

type StringStack struct {
	stackCnt   int
	stackSize  int
	stack      []*string
	stackIndex map[int]int
}

func NewStringStack(stackCnt, stackSize int) *StringStack {
	return &StringStack{
		stackCnt:   stackCnt,
		stackSize:  stackSize,
		stack:      make([]*string, stackSize*stackCnt),
		stackIndex: make(map[int]int),
	}
}

func (s *StringStack) push(stackNum int, element string) error {
	if s.isFull(stackNum) {
		return errors.New("the stack is full")
	}

	idx := s.stackIndex[stackNum]
	s.stack[(stackNum*s.stackSize)+idx] = &element
	s.stackIndex[stackNum]++
	return nil
}

func (s *StringStack) peek(stackNum int) (*string, error) {
	if s.isEmpty(stackNum) {
		return nil, errors.New("the stack is empty")
	}

	idx := s.stackIndex[stackNum]
	return s.stack[(stackNum*s.stackSize)+idx-1], nil
}

func (s *StringStack) pop(stackNum int) (*string, error) {
	if s.isEmpty(stackNum) {
		return nil, errors.New("the stack is empty")
	}

	idx := s.stackIndex[stackNum]
	data := s.stack[(stackNum*s.stackSize)+idx-1]
	s.stack[(stackNum*s.stackSize)+idx-1] = nil
	s.stackIndex[stackNum]--
	return data, nil
}

func (s *StringStack) isEmpty(stackNum int) bool {
	return s.stackIndex[stackNum] == 0
}

func (s *StringStack) isFull(stackNum int) bool {
	return s.stackIndex[stackNum] == s.stackSize-1
}

func main() {
	stackCnt := 3
	stackSize := 10
	s := NewStringStack(stackCnt, stackSize)
	testData := []string{"1", "2", "3", "4"}
	_ = s.push(0, testData[0])
	_ = s.push(1, testData[1])
	_ = s.push(2, testData[2])
	_ = s.push(2, testData[3])
	pe1, _ := s.peek(0)
	pe2, _ := s.peek(1)
	po1, _ := s.pop(2)
	po2, _ := s.pop(2)
	fmt.Println(s.stack, s.stackIndex)
	fmt.Println(*pe1, *pe2)
	fmt.Println(*po1, *po2)
}
