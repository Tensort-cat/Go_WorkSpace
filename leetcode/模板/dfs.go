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

import (
	"fmt"
)

var dir = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} // 四个方向

func dfs(grid [][]int, visited [][]bool, x, y int) {
	for i := 0; i < 4; i++ {
		nextx := x + dir[i][0]
		nexty := y + dir[i][1]
		if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
			continue // 越界了，直接跳过
		}
		if !visited[nextx][nexty] && grid[nextx][nexty] == 1 { // 没有访问过的 同时 是陆地的
			visited[nextx][nexty] = true
			dfs(grid, visited, nextx, nexty)
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n, n)
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
			if !visited[i][j] && grid[i][j] == 1 {
				visited[i][j] = true
				result++                 // 遇到没访问过的陆地，+1
				dfs(grid, visited, i, j) // 将与其链接的陆地都标记上 true
			}
		}
	}

	fmt.Println(result)
}
