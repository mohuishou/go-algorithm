package ShortestPath

import (
	"container/list"

	"errors"

	"github.com/mohuishou/algorithm/Graph"
)

//SPFADFS spfa算法dfs实现
func SPFADFS() {

}

//SPFABFS spfa算法bfs实现,负环判断不太稳定
func SPFABFS(graph Graph.Graph, s Graph.VextexType, dist []Graph.EdgeType, path []Graph.VextexType) error {
	//标记顶点是否在队列当中
	visited := make([]bool, graph.VNum)

	//统计，用于判断负环
	count := make([]int, graph.VNum)

	//初始化
	for i := 0; i < graph.VNum; i++ {
		dist[i] = INF      //所有的dist为无穷远，即为不可达
		path[i] = -1       //所有的上级节点为-1，即为无上级节点
		visited[i] = false // 所有都尚未观察
	}

	//先将源点入队
	dist[s] = 0
	path[s] = s
	visited[s] = true
	count[s] = 1

	q := list.New()
	q.PushBack(s)

	//循环跳出条件：队列为空
	for q.Len() != 0 {
		u := q.Front().Value.(Graph.VextexType)
		q.Remove(q.Front())

		//释放对点u的标记
		visited[u] = false

		e := graph.G[u].FisrtEdge

		for e != nil {
			//这条边下的顶点
			v := e.V

			//如果当前点的距离加上边的距离小于之前该点的距离，那么就更新该点的距离
			if dist[v] > dist[u]+e.Weight {
				dist[v] = dist[u] + e.Weight //更新该点距离
				path[v] = u                  //更新父节点

				//如果顶点不在队内，则将顶点入队
				if visited[v] == false {
					q.PushBack(v) //将该点入队
					visited[v] = true
					count[v]++

					//出现负环，报错
					if count[v] > graph.VNum {
						return errors.New("存在负环！")
					}
				}
			}
			e = e.Next
		}

	}
	return nil
}
