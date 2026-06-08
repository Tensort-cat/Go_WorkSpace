package main

import "fmt"

func isAnagram(s string, t string) bool {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	for _, ch := range t {
		cnt[ch-'a']--
	}

	for _, x := range cnt {
		if x != 0 {
			return false
		}
	}

	return true
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(isAnagram(s, t))
}
