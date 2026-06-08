package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0

	left, right := 0, 0
	hashMap := make(map[byte]int)
	for left <= right && right < len(s) {
		if _, exists := hashMap[s[right]]; !exists {
			hashMap[s[right]] = right
		} else {
			for ; left < hashMap[s[right]]+1; left++ {
				delete(hashMap, s[left])
			}
			hashMap[s[right]] = right
		}
		res = max(res, right-left+1)
		right++
	}

	return res
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}
