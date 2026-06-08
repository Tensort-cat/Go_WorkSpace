package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	res := false
	need := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if need == len(word) {
			res = true
			return
		}
		for i := 0; i < 4; i++ {
			nx, ny := x+dir[i][0], y+dir[i][1]
			if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) {
				continue
			}
			if !visited[nx][ny] && board[nx][ny] == word[need] {
				need++
				visited[nx][ny] = true
				dfs(nx, ny)
				visited[nx][ny] = false
				need--
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[need] {
				visited[i][j] = true
				need++
				dfs(i, j)
				visited[i][j] = false
				need--
			}
			if res {
				return res
			}
		}
	}

	return res
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	board := make([][]byte, m)
	var ch string
	for i := 0; i < m; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&ch)
			board[i][j] = ch[0]
		}
	}
	fmt.Println(board)

	var word string
	fmt.Scan(&word)

	fmt.Println(exist(board, word))
}
