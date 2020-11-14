package list

import (
	"container/list"
	"sync"
	"testing"
)

// TestList 测试标准库的链表
// go test -race .
func TestList(t *testing.T) {
	l := list.New()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			l.PushBack(i)
		}
		wg.Done()
	}()
	l.PushBack(11)
	wg.Wait()
}

func TestList_PushBack(t *testing.T) {
	l := New()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			l.PushBack(1)
		}
		wg.Done()
	}()
	l.PushBack(11)
	wg.Wait()
}
