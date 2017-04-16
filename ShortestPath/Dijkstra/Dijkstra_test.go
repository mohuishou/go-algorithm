package Diijkstra

import (
	"fmt"
	"testing"

	"github.com/mohuishou/algorithm/Graph"
)

func TestDijkstra(t *testing.T) {
	graph := Graph.BuildGraph("Dijkstra.txt")
	path := make([]Graph.VextexType, graph.VNum)
	dist := make([]Graph.EdgeType, graph.VNum)
	Dijkstra(graph, 0, dist, path)
	tPath, err := GetPathForDijkstra(path, 1)
	if err != nil {
		t.Fatal(err)
	}
	checkDist := []Graph.EdgeType{0, 7, 3, 9, 5}
	checkPath := []Graph.VextexType{1, 2, 0}

	fmt.Println("expected dist is \t\t", checkDist)
	fmt.Println("the value of dist is \t\t", dist)
	fmt.Println("expected path is \t\t", checkPath)
	fmt.Println("the value of path is \t\t", tPath)

	if len(checkPath) != len(tPath) || len(dist) != len(checkDist) {
		t.Fail()
	}

	for i := len(checkPath) - 1; i >= 0; i-- {
		if checkPath[i] != tPath[i] {
			t.Fail()
		}
	}

	for i := len(dist) - 1; i >= 0; i-- {
		if checkDist[i] != dist[i] {
			t.Fail()
		}
	}

}
