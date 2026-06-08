package main

// 牛顿迭代法
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	r := x
	for r > x/r {
		r = (r + x/r) / 2
	}

	return r
}
