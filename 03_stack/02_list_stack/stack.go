package stack

// node 节点
type node struct {
	prev, next *node
	value      string
}

// Stack 链式栈
type Stack struct {
	root *node
	len  int
}

// NewStack NewStack
func NewStack() *Stack {
	n := &node{}
	n.next = n
	n.prev = n
	return &Stack{root: n}
}

// Push 入栈
func (s *Stack) Push(item string) {
	n := &node{value: item}
	s.root.prev.next = n
	n.prev = s.root.prev
	n.next = s.root
	s.root.prev = n
	s.len++
}

// Pop 出栈
func (s *Stack) Pop() string {
	item := s.root.prev
	if item == s.root {
		return ""
	}

	s.root.prev = item.prev
	item.prev.next = s.root
	// 避免内存泄漏
	item.prev = nil
	item.next = nil
	s.len--
	return item.value
}
