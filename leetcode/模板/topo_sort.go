package main

// graph[i][j] = 1 表示 i -> j 有边
// 返回值：
//   - 若存在拓扑排序，返回排序结果
//   - 若不存在（图中有环），返回 nil

func topoSort(graph [][]int) []int {
	n := len(graph)

	// 统计入度
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] != 0 {
				indegree[j]++
			}
		}
	}

	// 初始化队列：入度为 0 的点入队
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 拓扑排序结果
	res := make([]int, 0)

	// BFS（Kahn 算法）
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		res = append(res, cur)

		// 删除 cur 的所有出边
		for next := 0; next < n; next++ {
			if graph[cur][next] != 0 {
				indegree[next]--

				if indegree[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	// 如果结果数量 != 节点数，说明有环
	if len(res) != n {
		return nil
	}

	return res
}
