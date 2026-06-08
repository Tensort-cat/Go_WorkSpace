package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	fmt.Println("v1:", v1, "v2:", v2)

	// 短的补0 (只有一个for循环会执行)
	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	for i := 0; i < len(v1); i++ {
		val1, _ := strconv.Atoi(v1[i])
		val2, _ := strconv.Atoi(v2[i])
		if val1 < val2 {
			return -1
		} else if val1 > val2 {
			return 1
		}
	}

	return 0
}

func main() {
	var version1, version2 string
	fmt.Scan(&version1, &version2)

	fmt.Printf("compareVersion(version1, version2): %v\n", compareVersion(version1, version2))
}
