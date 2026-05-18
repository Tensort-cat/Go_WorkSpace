package main

import "fmt"

/*
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
*/

/*
	组合问题：
	给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

	你可以按 任何顺序 返回答案。



	示例 1：

	输入：n = 4, k = 2
	输出：
	[
	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],
	]
	示例 2：

	输入：n = 1, k = 1
	输出：[[1]]


	提示：

	1 <= n <= 20
	1 <= k <= n
*/
func combine(n int, k int) [][]int {
	path, res := make([]int, 0, k), make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k { // 说明已经满足了k个数的要求
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
			if n-i+1 < k-len(path) { // 剪枝
				break
			}
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return res
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(combine(n, k))
}
