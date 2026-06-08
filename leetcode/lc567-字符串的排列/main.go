package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	cnt1 := [26]int{}
	cnt2 := [26]int{}
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}

	left, right := 0, len(s1)
	for right < len(s2) {
		cnt2[s2[left]-'a']--
		cnt2[s2[right]-'a']++
		if cnt1 == cnt2 {
			return true
		}
		left++
		right++
	}

	return false
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Printf("checkInclusion(s1, s2): %v\n", checkInclusion(s1, s2))
}
