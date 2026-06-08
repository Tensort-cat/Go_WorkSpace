package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := 0
	path := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			path = append(path, root.Val)
			if root.Left == nil && root.Right == nil { // 叶子
				num := 0
				for _, v := range path {
					num = num*10 + v
				}
				res += num
			}
			dfs(root.Left)
			dfs(root.Right)
			path = path[:len(path)-1]
		}
	}

	dfs(root)
	return res
}
