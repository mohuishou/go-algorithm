package GraphMatrix

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := BuildGraph("test.txt")
	fmt.Println(graph)
}
