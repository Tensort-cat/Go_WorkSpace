package main

import (
	"fmt"
	"sort"
)

// 暴力解
func topKFrequent1(nums []int, k int) []int {
	res := []int{}
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	type st struct {
		num int
		cnt int
	}
	sts := []st{}
	for key, value := range hashMap {
		sts = append(sts, st{num: key, cnt: value})
	}

	sort.Slice(sts, func(i, j int) bool {
		return sts[i].cnt > sts[j].cnt
	})

	for i := 0; i < k; i++ {
		res = append(res, sts[i].num)
	}

	return res
}

// 堆
func topKFrequent2(nums []int, k int) []int {
	// 第一步：统计每个元素的出现次数
	cnt := map[int]int{}
	maxCnt := 0
	for _, x := range nums {
		cnt[x]++
		maxCnt = max(maxCnt, cnt[x])
	}

	// 第二步：把出现次数相同的元素，放到同一个桶中
	buckets := make([][]int, maxCnt+1)
	for x, c := range cnt {
		buckets[c] = append(buckets[c], x)
	}

	// 第三步：倒序遍历 buckets，把出现次数前 k 大的元素加入答案
	ans := []int{} // 预分配空间
	// 注意题目保证答案唯一，一定会出现某次 append 后 len(ans) 恰好等于 k 的情况
	for i := maxCnt; len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}

func main() {
	var k, n int
	fmt.Scan(&n)
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Scan(&k)

	fmt.Println(topKFrequent2(nums, k))
}
