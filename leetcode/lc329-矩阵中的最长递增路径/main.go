package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(matrix [][]int, row, col, length int)
	dfs = func(matrix [][]int, row, col, length int) {
		memo[row][col] = 1
		res = max(res, length)
		if row-1 >= 0 && matrix[row-1][col] > matrix[row][col] { // 向上
			dfs(matrix, row-1, col, length+1)
		}
		if row+1 < m && matrix[row+1][col] > matrix[row][col] { // 向下
			dfs(matrix, row+1, col, length+1)
		}
		if col-1 >= 0 && matrix[row][col-1] > matrix[row][col] { // 向左
			dfs(matrix, row, col-1, length+1)
		}
		if col+1 < n && matrix[row][col+1] > matrix[row][col] { // 向右
			dfs(matrix, row, col+1, length+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if memo[i][j] != 1 {
				dfs(matrix, i, j, 0)
			}
		}
	}

	return res + 1
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	matrix := make([][]int, m, m)
	var elem int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&elem)
			matrix[i] = append(matrix[i], elem)
		}
	}
	fmt.Println(matrix)

	fmt.Println(longestIncreasingPath(matrix))
}
