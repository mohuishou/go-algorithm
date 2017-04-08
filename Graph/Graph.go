package Graph

// EdgeType 边的权值类型
type EdgeType int

// VextexType 顶点类型定义
type VextexType int

// VextexDataType 顶点值类型定义
type VextexDataType int

//EdgeNode 边的节点
type EdgeNode struct {
	weight EdgeType   //权值
	v      VextexType //指向储存该顶点的下标
	next   *EdgeNode  //指向下一条边
}

//VextexNode 顶点节点定义
type VextexNode struct {
	data      VextexDataType //顶点的值
	FisrtEdge *EdgeNode      //该顶点指向的第一条边
}

//Graph 图
type Graph []VextexNode

//CreateGraph 创建邻接表
func CreateGraph(n int) (graph Graph) {
	graph = make(Graph, n)
	for i := 0; i < n; i++ {
		graph[i] = VextexNode{}
	}
	return graph
}

//AddEdge 添加边
func (graph Graph) AddEdge(s, t VextexType, weight EdgeType) {
	edge := &EdgeNode{v: t, weight: weight}

	//添加边到头部
	edge.next = graph[s].FisrtEdge
	graph[s].FisrtEdge = edge
}
