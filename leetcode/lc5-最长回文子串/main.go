package main

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isVaild(s, i, j) && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	return res
}

func isVaild(s string, i, j int) bool {
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
