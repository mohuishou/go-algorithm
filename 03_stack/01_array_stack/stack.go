package stack

// Stack Stack
type Stack struct {
	items   []string
	current int
}

// NewStack NewStack
func NewStack() *Stack {
	return &Stack{
		items:   make([]string, 10),
		current: 0,
	}
}

// Push 入栈
func (s *Stack) Push(item string) {
	s.current++
	// 判断底层 slice 是否满了，如果满了就 append
	if s.current == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	s.items[s.current] = item
}

// Pop 出栈
func (s *Stack) Pop() string {
	if s.current == 0 {
		return ""
	}
	item := s.items[s.current]
	s.current--
	return item
}
