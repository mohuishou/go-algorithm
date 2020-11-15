package lru

import (
	"fmt"
)

// Node 链表的节点
type Node struct {
	prev, next *Node

	list *LRU

	key   string
	value interface{}
}

// LRU 缓存
type LRU struct {
	root *Node // 根节点
	cap  int   // 当前缓存容量
	len  int   // 缓存的长度
}

// NewLRU NewLRU
func NewLRU(cap int) *LRU {
	l := &LRU{
		root: &Node{},
		cap:  cap,
	}
	l.root.prev = l.root
	l.root.next = l.root
	l.root.list = l
	return l
}

// Get 获取缓存数据
// 如果获取到数据，就把这个节点移动到链表头部
// 如果没有获取到，就返回nil
func (l *LRU) Get(key string) interface{} {
	defer l.debug()
	n := l.get(key)
	if n == nil {
		return nil
	}

	return n.value
}

func (l *LRU) get(key string) *Node {
	for n := l.root.next; n != l.root; n = n.next {
		if n.key == key {
			n.prev.next = n.next
			n.next.prev = n.prev

			n.next = l.root.next
			l.root.next.prev = n
			l.root.next = n
			n.prev = l.root
			return n
		}
	}
	return nil
}

// Put 写入缓存数据
// 如果 key 已经存在，那么更新值
// 如果 key 不存在，那么插入到第一个节点
// 当缓存容量满了的时候，会自动删除最后的数据
func (l *LRU) Put(key string, value interface{}) {
	defer l.debug()
	n := l.get(key)
	if n != nil {
		n.value = value
		return
	}

	// 缓存满了
	if l.len == l.cap {
		last := l.root.prev
		last.prev.next = l.root
		l.root.prev = last.prev
		last.list = nil
		last.prev = nil
		last.next = nil
		l.len--
	}

	node := &Node{key: key, value: value}
	head := l.root.next
	head.prev = node
	node.next = head
	node.prev = l.root
	l.root.next = node
	l.len++
	node.list = l
}

// debug for debug
func (l *LRU) debug() {
	fmt.Println("lru len: ", l.len)
	fmt.Println("lru cap: ", l.cap)
	for n := l.root.next; n != l.root; n = n.next {
		fmt.Printf("%s:%v -> ", n.key, n.value)
	}
	fmt.Println()
	fmt.Println()
}
