package ShortestPath

import (
	"testing"

	"fmt"

	"github.com/mohuishou/algorithm/Graph"
	"github.com/mohuishou/algorithm/GraphMatrix"
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

func TestFolyd(t *testing.T) {
	graph := GraphMatrix.BuildGraph("Floyd.txt")
	dist := make([][]GraphMatrix.EdgeType, graph.VNum)
	path := make([][]int, graph.VNum)
	for i := 0; i < graph.VNum; i++ {
		path[i] = make([]int, graph.VNum)
		dist[i] = make([]GraphMatrix.EdgeType, graph.VNum)
	}
	Floyd(graph, dist, path)

	fmt.Println("test result：")
	fmt.Print("source -> dest \t")
	fmt.Print("dist \t\t")
	fmt.Print("path \t\t\n")

	for i := 0; i < graph.VNum; i++ {
		for j := 0; j < graph.VNum; j++ {
			if i == j {
				continue
			}
			fmt.Print(i, " -> ", j, " \t\t")
			fmt.Print(dist[i][j], " \t\t")
			tPath := GetPathForFloyd(path, i, j)
			for i := 0; i < len(tPath); i++ {
				fmt.Print(tPath[i])
				if i < len(tPath)-1 {
					fmt.Print("->")
				}
			}
			fmt.Print("\t\t\n")
		}
	}
}

func TestSPFABFS(t *testing.T) {
	graph := Graph.BuildGraph("Dijkstra.txt")
	path := make([]Graph.VextexType, graph.VNum)
	dist := make([]Graph.EdgeType, graph.VNum)
	err := SPFABFS(graph, 0, dist, path)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dist)
}
