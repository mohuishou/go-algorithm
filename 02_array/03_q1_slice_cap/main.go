// Package main What will eventually be output?
// A: 5 8    B: 8 8   C: 5 5   D: 5 6
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Println(len(s), cap(s))
}
