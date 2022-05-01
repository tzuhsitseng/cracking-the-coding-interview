package shared

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

//

type TreeNodeStack struct {
	stack []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{
		stack: make([]*TreeNode, 0),
	}
}

func (s *TreeNodeStack) push(element *TreeNode) {
	if s.stack == nil {
		s.stack = make([]*TreeNode, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *TreeNodeStack) pop() *TreeNode {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return result
}

func (s *TreeNodeStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *TreeNodeStack) peek() *TreeNode {
	if s.isEmpty() {
		return nil
	}

	return s.stack[len(s.stack)-1]
}

//

type StringStack struct {
	stack []string
}

func NewStringStack() *StringStack {
	return &StringStack{
		stack: make([]string, 0),
	}
}

func (s *StringStack) push(element string) {
	if s.stack == nil {
		s.stack = make([]string, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *StringStack) pop() *string {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return &result
}

func (s *StringStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *StringStack) peek() *string {
	if s.isEmpty() {
		return nil
	}

	return &s.stack[len(s.stack)-1]
}
