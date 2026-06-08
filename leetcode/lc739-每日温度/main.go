package main

// 暴力
func dailyTemperatures1(temperatures []int) []int {
	res := []int{}

	for i := 0; i < len(temperatures)-1; i++ {
		have := false
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res = append(res, j-i)
				have = true
				break
			}
		}
		if !have {
			res = append(res, 0)
		}
	}

	res = append(res, 0)
	return res
}

// 	单调栈
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	stack := []int{} // 存下标

	for i := 0; i < n; i++ {

		// 当前温度更高
		for len(stack) > 0 &&
			temperatures[i] > temperatures[stack[len(stack)-1]] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top] = i - top
		}

		stack = append(stack, i)
	}

	return res
}
