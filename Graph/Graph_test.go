package Graph

import (
    "fmt"
    "testing"
)

func TestGraph(t *testing.T) {
    graph := BuildGraph("test.txt")
    fmt.Println(graph)
    e := graph[1].FisrtEdge
    for e != nil {
        if e.v != 2 || e.weight != 3 {
            t.Fatal("错误：建图失败，e.v!=2,e.weight!=3")
        }
        e = e.next
    }
}
