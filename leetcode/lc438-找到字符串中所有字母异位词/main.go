package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	res := []int{}
	n, m := len(s), len(p)
	if n < m {
		return res
	}

	var sCnt, pCnt [26]int

	// 统计 p 和 s 的第一个窗口
	for i := 0; i < m; i++ {
		pCnt[p[i]-'a']++
		sCnt[s[i]-'a']++
	}

	// 判断第一个窗口
	if sCnt == pCnt {
		res = append(res, 0)
	}

	// 滑动窗口
	for right := m; right < n; right++ {
		left := right - m

		sCnt[s[left]-'a']--  // 移出左边字符
		sCnt[s[right]-'a']++ // 加入右边字符

		if sCnt == pCnt {
			res = append(res, left+1)
		}
	}

	return res
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
