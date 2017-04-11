//Package GraphMatrix 邻接矩阵表示图
package GraphMatrix

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// EdgeType 边的权值类型
type EdgeType int

//Graph 图
type Graph struct {
	VNum, ENum int          //顶点、边的个数
	G          [][]EdgeType //邻接矩阵
}

//CreateGraph 创建一个图并且初始化邻接矩阵
func CreateGraph(n int) (graph Graph) {
	graph = Graph{VNum: n}
	graph.G = make([][]EdgeType, n)
	for i := 0; i < n; i++ {
		graph.G[i] = make([]EdgeType, n)
	}
	return graph
}

//BuildGraph 创建图
func BuildGraph(path string) (graph Graph) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	i := 0
	//边的数目
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return graph
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		data := strings.Split(line, " ")
		if i == 0 {
			n, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			graph = CreateGraph(n)

			graph.ENum, err = strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}
		} else if i <= graph.ENum {
			s, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			t, err := strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}
			weight, err := strconv.Atoi(data[2])
			if err != nil {
				panic(err)
			}
			graph.G[s][t] = EdgeType(weight)
		}
		i++
	}

}
