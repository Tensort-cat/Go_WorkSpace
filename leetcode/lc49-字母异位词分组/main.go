package main

import (
	"fmt"
)

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	l := len(strs)
	tag := make([]bool, l)
	for i, str := range strs {
		if tag[i] {
			continue
		}
		m1 := make(map[rune]int)
		for _, r := range str {
			m1[r]++
		}

		var strs_ []string
		strs_ = append(strs_, str)

		for j := i + 1; j < l; j++ {
			isAppend := true
			m2 := make(map[rune]int)
			for _, r := range strs[j] {
				if _, exists := m1[r]; !exists {
					isAppend = false
					break
				}
				m2[r]++
			}
			if !isAppend {
				continue
			}
			for k, _ := range m1 {
				if m1[k] != m2[k] {
					isAppend = false
					break
				}
			}
			if !isAppend {
				continue
			}
			strs_ = append(strs_, strs[j])
			tag[j] = true
		}
		res = append(res, strs_)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
