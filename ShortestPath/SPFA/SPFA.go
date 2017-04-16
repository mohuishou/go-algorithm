package SPFA

import (
	"container/list"

	"errors"

	"github.com/mohuishou/algorithm/Graph"
)

//INF 定义一个无穷大的数
const INF = 0xfffff

var (
	graph   Graph.Graph        //图
	s       Graph.VextexType   //需要查找的顶点
	dist    []Graph.EdgeType   //各顶点到s的最短路径
	path    []Graph.VextexType //路径集合
	visited []bool             //标记顶点是否在队列当中
	count   []int              //统计，用于判断负环
	isInit  bool               //是否初始化
)

//Init 初始化
func Init(g Graph.Graph, S Graph.VextexType, Dist []Graph.EdgeType, Path []Graph.VextexType) {
	graph = g
	s = S
	dist = Dist
	path = Path
	//标记顶点是否在队列当中
	visited = make([]bool, graph.VNum)

	//统计，用于判断负环
	count = make([]int, graph.VNum)

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

	isInit = true
}

//DFS spfa算法dfs实现
func DFS(u Graph.VextexType) error {
	if isInit != true {
		return errors.New("请先执行SPFA.Init方法")
	}
	visited[u] = true
	e := graph.G[u].FisrtEdge
	for e != nil {
		v := e.V
		if dist[v] > dist[u]+e.Weight {
			dist[v] = dist[u] + e.Weight //更新该点距离
			path[v] = u                  //更新父节点
			if visited[v] == false {
				count[v]++
				if count[v] > graph.VNum {
					return errors.New("存在负环！")
				}

				//注意DFS的结果不能直接return，直接return的时候回溯的时候就没有办法在上一级重新找值了
				err := DFS(v)
				if err != nil {
					return err
				}
			} else {
				return nil
			}
		}
		e = e.Next
	}
	visited[u] = false
	return nil
}

//BFS spfa算法bfs实现,负环判断不太稳定
func BFS() error {
	if isInit != true {
		return errors.New("请先执行SPFA.Init方法")
	}

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
