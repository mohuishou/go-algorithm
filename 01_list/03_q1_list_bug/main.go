package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.List{}
	e := l.PushBack(10)

	l = list.List{}
	l.Remove(e)
	fmt.Println("list len: ", l.Len())
}
