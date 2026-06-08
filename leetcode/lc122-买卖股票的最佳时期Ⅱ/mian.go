package main

// 贪心
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}

	return res
}
