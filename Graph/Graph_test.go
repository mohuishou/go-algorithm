package Graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := CreateGraph(3)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 3)
}
