package shared

type ListNode struct {
	Val  int
	Next *ListNode
}

func push(stack []*ListNode, elements ...*ListNode) []*ListNode {
	return append(stack, elements...)
}

func pop(stack []*ListNode) (*ListNode, []*ListNode) {
	l := len(stack)
	if l > 0 {
		return stack[l-1], stack[:l-1]
	}
	return nil, nil
}
