package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int
	l := len(intervals)
	if l == 1 {
		res = append(res, intervals[0])
		return res
	}

	// 先按区间起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < l; i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
			intervals[i][1] = max(intervals[i][1], intervals[i-1][1])
		} else {
			res = append(res, intervals[i-1])
		}
	}
	res = append(res, intervals[l-1])

	return res
}

func main() {
	fmt.Printf("merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}): %v\n", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Printf("merge([][]int{{1, 4}, {4, 5}}): %v\n", merge([][]int{{1, 4}, {4, 5}}))
	fmt.Printf("merge([][]int{{1, 4}, {0, 4}}): %v\n", merge([][]int{{1, 4}, {0, 4}}))
	fmt.Printf("merge([][]int{{1, 4}, {2, 3}}): %v\n", merge([][]int{{1, 4}, {2, 3}}))
	fmt.Printf("merge([][]int{{1, 4}, {0, 0}}): %v\n", merge([][]int{{1, 4}, {0, 0}}))
}
