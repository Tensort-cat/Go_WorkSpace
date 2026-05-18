package main

import "fmt"

func main() {
	var numCourses, n int
	fmt.Scan(&numCourses, &n)
	prerequisites := make([][]int, n, n)
	graph := [2000][2000]int{}
	for i := 0; i < n; i++ {
		prerequisites[i] = make([]int, 2)
		fmt.Scan(&prerequisites[i][0], &prerequisites[i][1])
		graph[prerequisites[i][1]][prerequisites[i][0]] = 1
	}

	// 求 graph 的拓扑排序是否存在即可
	topo := topoSort(graph)
	fmt.Println(topo != nil)
}

func topoSort(graph [2000][2000]int) []int {
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
