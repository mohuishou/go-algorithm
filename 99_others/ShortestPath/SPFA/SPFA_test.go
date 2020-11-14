package SPFA

import (
	"fmt"
	"testing"

	"github.com/mohuishou/algorithm/Graph"
)

func TestBFS(t *testing.T) {
	graph := Graph.BuildGraph("spfa.txt")
	path := make([]Graph.VextexType, graph.VNum)
	dist := make([]Graph.EdgeType, graph.VNum)
	Init(graph, 0, dist, path)
	err := BFS()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dist)
}

func TestDFS(t *testing.T) {
	graph := Graph.BuildGraph("spfa.txt")
	path := make([]Graph.VextexType, graph.VNum)
	dist := make([]Graph.EdgeType, graph.VNum)
	Init(graph, 0, dist, path)
	err := DFS(s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dist)
}
