package main

/*
99. 计数孤岛
题目描述
给定一个由 1（陆地）和 0（水）组成的矩阵，你需要计算岛屿的数量。岛屿由水平方向或垂直方向上相邻的陆地连接而成，并且四周都是水域。你可以假设矩阵外均被水包围。

输入描述
第一行包含两个整数 N, M，表示矩阵的行数和列数。

后续 N 行，每行包含 M 个数字，数字为 1 或者 0。

输出描述
输出一个整数，表示岛屿的数量。如果不存在岛屿，则输出 0。
输入示例
4 5
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1
输出示例
3
*/

import "fmt"

var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

// bfs 广度优先搜索
func bfs(grid [][]int, visited [][]bool, x, y int) {
	// 队列中存储坐标
	queue := [][2]int{{x, y}}

	for len(queue) > 0 {
		// 取出队头元素
		cur := queue[0]
		queue = queue[1:]

		curX := cur[0]
		curY := cur[1]

		// 向四个方向扩展
		for i := 0; i < 4; i++ {
			nextX := curX + dir[i][0]
			nextY := curY + dir[i][1]

			// 越界
			if nextX < 0 || nextX >= len(grid) ||
				nextY < 0 || nextY >= len(grid[0]) {
				continue
			}

			// 没访问过且是陆地
			if !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
				visited[nextX][nextY] = true
				queue = append(queue, [2]int{nextX, nextY})
			}
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			// 找到一个新的岛屿
			if !visited[i][j] && grid[i][j] == 1 {
				visited[i][j] = true
				result++

				// bfs 把整个岛屿遍历完
				bfs(grid, visited, i, j)
			}
		}
	}

	fmt.Println(result)
}
