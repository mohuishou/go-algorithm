package ShortestPath

import (
	"testing"

	"fmt"

	"github.com/mohuishou/algorithm/Graph"
)

func TestDijkstra(t *testing.T) {
	graph := Graph.BuildGraph("Dijkstra.txt")
	path := make([]Graph.VextexType, graph.VNum)
	dist := make([]Graph.EdgeType, graph.VNum)
	Dijkstra(graph, 1, dist, path)
	fmt.Println(dist)
}
