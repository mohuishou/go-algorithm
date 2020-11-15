package list

import (
	"container/list"
	"sync"
)

// List 链表
type List struct {
	*list.List
	mu sync.Mutex
}

// New 新建链表
func New() *List {
	return &List{List: list.New()}
}

// PushBack 像链表尾部插入值
func (l *List) PushBack(v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.List.PushBack(v)
}
