//Package Graph 邻接表
package Graph

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// EdgeType 边的权值类型
type EdgeType int

// VextexType 顶点类型定义
type VextexType int

// VextexDataType 顶点值类型定义
type VextexDataType int

//EdgeNode 边的节点
type EdgeNode struct {
	Weight EdgeType   //权值
	V      VextexType //指向储存该顶点的下标
	Next   *EdgeNode  //指向下一条边
}

//VextexNode 顶点节点定义
type VextexNode struct {
	data      VextexDataType //顶点的值
	FisrtEdge *EdgeNode      //该顶点指向的第一条边
}

//Graph 图
type Graph struct {
	VNum, ENum int          //顶点数目，边数目
	G          []VextexNode //邻接表
}

//CreateGraph 创建邻接表
func CreateGraph(VNum int) (graph Graph) {
	graph.VNum = VNum
	graph.G = make([]VextexNode, VNum)
	for i := 0; i < VNum; i++ {
		graph.G[i] = VextexNode{}
	}
	return graph
}

//AddEdge 添加边
func (graph Graph) AddEdge(s, t VextexType, weight EdgeType) {
	edge := &EdgeNode{V: t, Weight: weight}

	//添加边到头部
	edge.Next = graph.G[s].FisrtEdge
	graph.G[s].FisrtEdge = edge
}

//BuildGraph 通过读取文件建图
//文件格式要求:
//顶点个数 边数
//顶点v1 顶点V2 边的权重
//...
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
			graph.AddEdge(VextexType(s), VextexType(t), EdgeType(weight))
		}
		i++
	}

}
