package main

import (
	"fmt"
	"log"
)

func main() {
	// 初始化数组
	var arr = [3]int{1, 2, 3}
	// 查找元素
	fmt.Printf("arr[1]: %d\n", arr[1])
	// 删除元素
	remove(&arr, 2)
	// remove(&arr, 3) // will panic
	fmt.Println(arr)

	fmt.Printf("main: %p --> %+v\n", &arr, arr)
	p(arr)
	p2(&arr)
}

// 删除数组 arr 的某个元素
// index 为需要删除的索引
// 从需要删除的元素开始，依次将后面的元素往前移动一位即可
// 然后将最后一位修改为该类型的默认值
func remove(arr *[3]int, index int) {
	if index >= len(arr) {
		log.Panicf("%d remove out range arr", index)
	}
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = 0
}

func p(arr [3]int) {
	fmt.Printf("p: %p --> %+v\n", &arr, arr)
}

func p2(arr *[3]int) {
	fmt.Printf("p2: %p --> %+v\n", arr, arr)
}
