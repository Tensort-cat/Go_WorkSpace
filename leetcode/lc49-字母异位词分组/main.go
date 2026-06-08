package main

import (
	"fmt"
)

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	res := make([][]string, 0, len(mp))
	for _, group := range mp {
		res = append(res, group)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
