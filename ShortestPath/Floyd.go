package ShortestPath

import (
	"errors"

	"github.com/mohuishou/algorithm/GraphMatrix"
)

//Floyd 求取多源最短路径
func Floyd(graph GraphMatrix.Graph, dist [][]GraphMatrix.EdgeType, path [][]int) error {
	for i := 0; i < graph.VNum; i++ {
		for j := 0; j < graph.VNum; j++ {
			path[i][j] = -1
		}
	}

	//使dist等于邻接矩阵
	dist = graph.G

	for k := 0; k < graph.VNum; k++ {
		for i := 0; i < graph.VNum; i++ {
			for j := 0; j < graph.VNum; j++ {

				//找到更短的路径
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]

					//发现负值圈
					if i == j && dist[i][j] < 0 {
						return errors.New("存在负值圈")
					}
					path[i][j] = k
				}
			}
		}
	}
	return nil
}
