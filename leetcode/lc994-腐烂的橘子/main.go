package main

import "fmt"

func main() {
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var m, n int
	fmt.Scan(&m, &n)

	grid := make([][]int, m, m)
	queue := [][2]int{}
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
			if grid[i][j] == 2 { // 是烂橘子
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	res := 0
	// 对烂橘子逐层bfs
	for {
		nextQueue := [][2]int{}
		isRotting := false
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				nextX := cur[0] + dir[i][0]
				nextY := cur[1] + dir[i][1]
				if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
					continue
				}
				if grid[nextX][nextY] == 1 { // 好橘子才过去
					grid[nextX][nextY] = 2 // 污染好橘子
					isRotting = true       // 发生了污染
					nextQueue = append(nextQueue, [2]int{nextX, nextY})
				}
			}
		}
		if isRotting {
			res++
		}
		if len(nextQueue) == 0 {
			break
		}
		queue = nextQueue
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = -1
			}
		}
	}
	fmt.Println(res)
	fmt.Println(grid)
}
