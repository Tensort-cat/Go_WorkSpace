package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var record_row, record_col []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				record_row = append(record_row, i)
				record_col = append(record_col, j)
			}
		}
	}

	cnt := len(record_row)
	for i := 0; i < cnt; i++ {
		toZero(matrix, record_row[i], record_col[i], m, n)
	}
}

func toZero(matrix [][]int, row int, col int, m int, n int) {
	for i := 0; i < n; i++ {
		matrix[row][i] = 0
	}

	for i := 0; i < m; i++ {
		matrix[i][col] = 0
	}
}
