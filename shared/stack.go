package shared

type StackOfListNode struct {
	stack []*ListNode
}

func (s *StackOfListNode) push(elements ...*ListNode) {
	if s.stack == nil {
		s.stack = make([]*ListNode, 0)
	}
	s.stack = append(s.stack, elements...)
}

func (s *StackOfListNode) pop() *ListNode {
	l := len(s.stack)
	if l == 0 {
		return nil
	}

	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return result
}
