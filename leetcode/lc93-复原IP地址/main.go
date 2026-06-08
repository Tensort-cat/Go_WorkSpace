package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}
	var dfs func(int)
	dfs = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		for i := start; i < len(s); i++ {
			// 有前导零
			if i != start && s[start] == '0' {
				break
			}
			str := s[start : i+1]
			val, _ := strconv.Atoi(str)
			if val >= 0 && val <= 255 {
				path = append(path, str)
				dfs(i + 1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}

	dfs(0)
	return res
}
