package shared

type ListNodeStack struct {
	stack []*ListNode
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
