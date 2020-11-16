package main

import "fmt"

func main() {
	// 初始化
	s1 := make([]int, 2)
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	// s1(0xc00000c080): [0 0], len: 2, cap: 2

	// 赋值
	s1[0] = 1
	s1[1] = 2
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	//	s1(0xc00000c080): [1 2], len: 2, cap: 2

	// 扩容
	s1 = append(s1, 3, 4, 5, 6, 7)
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	// s1(0xc00000c080): [1 2 3], len: 3, cap: 4

	// 删除元素
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	// s1(0xc00000c080): [1 3], len: 2, cap: 4

	// 复制一个 slice
	s2 := s1[:2]
	fmt.Printf("s2(%p): %v, len: %d, cap: %d\n", &s2, s2, len(s2), cap(s2))
	// s2(0xc00000c120): [1 3], len: 2, cap: 4

	s1[0] = 10 // 这里可以发现，s1[0] s2[0] 都被修改为了 10
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	// s1(0xc00000c080): [10 3], len: 2, cap: 4
	fmt.Printf("s2(%p): %v, len: %d, cap: %d\n", &s2, s2, len(s2), cap(s2))
	// s2(0xc00000c120): [10 3], len: 2, cap: 4

	s1 = append(s1, 5, 6, 7, 8)
	s1[0] = 11 // 这里可以发现，s1[0] 被修改为了 11, s2[0] 还是10
	fmt.Printf("s1(%p): %v, len: %d, cap: %d\n", &s1, s1, len(s1), cap(s1))
	// s1(0xc00011c020): [11 3 5 6 7 8], len: 6, cap: 8
	fmt.Printf("s2(%p): %v, len: %d, cap: %d\n", &s2, s2, len(s2), cap(s2))
	// s2(0xc00011c0c0): [10 3], len: 2, cap: 4

	// 对比一下两种空 slice 的区别
	var s3 []int
	s4 := make([]int, 0)
	fmt.Printf("%p --> %#v\n", s3, s3) // 0x0 --> []int(nil)
	fmt.Printf("%p --> %#v\n", s4, s4) // 0x587450 --> []int{}

	initSlice()
}

func sliceChange(s []int) {
	// 不允许直接这么操作
	s[0] = 1
}

func initSlice() {
	var s1 []int
	fmt.Printf("%p: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))
	s2 := make([]uint, 0)
	fmt.Printf("%p: %v, len: %d, cap: %d\n", s2, s2, len(s2), cap(s2))
	s3 := []int{}
	fmt.Printf("%p: %v, len: %d, cap: %d\n", s3, s3, len(s3), cap(s3))
}
