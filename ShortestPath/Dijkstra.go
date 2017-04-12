package ShortestPath

import (
	"errors"

	"container/list"

	"github.com/mohuishou/algorithm/Graph"
)

//INF 无穷大
const INF = 0xffffff

//Dijkstra 算法
//一种求单源最短路径的算法
func Dijkstra(graph Graph.Graph, s Graph.VextexType, dist []Graph.EdgeType, path []Graph.VextexType) {
	visited := make([]bool, graph.VNum)
	//初始化
	for i := 0; i < graph.VNum; i++ {
		dist[i] = INF //距离为无穷大
		path[i] = -1  //没有上一个节点
		visited[i] = false
	}
	path[s] = s
	dist[s] = 0

	//使用list实现一个队列操作
	q := list.New()

	//将点s入队
	q.PushBack(s)
	for q.Len() != 0 {
		u := q.Front().Value.(Graph.VextexType)
		q.Remove(q.Front())
		//如果该点周围的点已经走过，则无需再走
		if visited[u] {
			continue
		}

		//将该点加入已观察
		visited[u] = true

		e := graph.G[u].FisrtEdge

		for e != nil {
			//这条边下的顶点
			v := e.V

			//如果该点尚未走过，并且当前点的距离加上边的距离小于之前该点的距离，那么就更新该点的距离
			if visited[v] == false && dist[v] > dist[u]+e.Weight {
				dist[v] = dist[u] + e.Weight //更新该点距离
				path[v] = u                  //更新父节点
				q.PushBack(v)                //将该点入队
			}
			e = e.Next
		}

	}

}

//GetPathForDijkstra 通过路径获得到指定目的节点的路径
func GetPathForDijkstra(path []Graph.VextexType, t Graph.VextexType) ([]Graph.VextexType, error) {
	tPath := make([]Graph.VextexType, 0)
	for {
		tPath = append(tPath, t)
		if path[t] == -1 {
			return nil, errors.New("不存在到该节点的路径")
		}
		if t == path[t] {
			return tPath, nil
		}
		t = path[t]
	}
}
