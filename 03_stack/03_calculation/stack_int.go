package calculation

// StackInt StackInt
type StackInt struct {
	items   []int
	current int
}

// NewStackInt NewStackInt
func NewStackInt() *StackInt {
	return &StackInt{
		items:   make([]int, 10),
		current: 0,
	}
}

// Push 入栈
func (s *StackInt) Push(item int) {
	s.current++
	// 判断底层 slice 是否满了，如果满了就 append
	if s.current == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	s.items[s.current] = item
}

// Pop 出栈
func (s *StackInt) Pop() int {
	if s.current == 0 {
		return 0
	}
	item := s.items[s.current]
	s.current--
	return item
}
