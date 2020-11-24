package main

import "fmt"

func main() {
	f0()
	f1()
	f2()

	fmt.Println("f3:", f3())
	fmt.Println("f4:", f4())

	f5()
}

// 基本用法：延迟调用，清理资源
func f0() {
	defer fmt.Println("clean")
	fmt.Println("hello")
}

// 基本用法1: 后进先出
func f1() {
	defer fmt.Println("1")
	defer fmt.Println("2")

	fmt.Println("3")
}

// 基本用法2：异常恢复
func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("paniced: %+v", err)
		}
	}()
	panic("test")
}

// 容易掉坑1：函数变量修改
func f3() (res int) {
	defer func() {
		res++
	}()
	return 0
}

// 容易掉坑之2：参数复制
func f4() (res int) {
	defer func(res int) {
		res++
	}(res)
	return 0
}

// 容易掉坑3：值
func f5() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 循环defer不可取
func f6() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("f6: %d\n", i)
		}()
	}
}
