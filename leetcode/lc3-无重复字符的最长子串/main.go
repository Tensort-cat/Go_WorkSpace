package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if l-i <= res {
			break
		}
		count := 0
		m := make(map[byte]struct{})
		for j := i; j < l; j++ {
			_, exists := m[s[j]]
			if exists {
				break
			}
			m[s[j]] = struct{}{}
			count++
		}
		res = max(count, res)
	}

	return res
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}
