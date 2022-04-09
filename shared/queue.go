package shared

type ListNodeQueue struct {
	queue []*ListNode
}

func NewListNodeQueue() *ListNodeQueue {
	return &ListNodeQueue{
		queue: make([]*ListNode, 0),
	}
}

func (q *ListNodeQueue) enqueue(element *ListNode) {
	if q.queue == nil {
		q.queue = make([]*ListNode, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *ListNodeQueue) dequeue() *ListNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *ListNodeQueue) peek() *ListNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *ListNodeQueue) isEmpty() bool {
	return len(q.queue) == 0
}

//

type TreeNodeQueue struct {
	queue []*TreeNode
}

func NewTreeNodeQueue() *TreeNodeQueue {
	return &TreeNodeQueue{
		queue: make([]*TreeNode, 0),
	}
}

func (q *TreeNodeQueue) enqueue(element *TreeNode) {
	if q.queue == nil {
		q.queue = make([]*TreeNode, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *TreeNodeQueue) dequeue() *TreeNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *TreeNodeQueue) peek() *TreeNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *TreeNodeQueue) isEmpty() bool {
	return len(q.queue) == 0
}
