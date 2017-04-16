package ShortestPath

import (
	"testing"

	"fmt"

	"github.com/mohuishou/algorithm/GraphMatrix"
)

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
