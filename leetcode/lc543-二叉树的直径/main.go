package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lLen := dfs(node.Left)
		rLen := dfs(node.Right)
		ans = max(ans, lLen+rLen) // 两条链拼成路径
		return max(lLen, rLen) + 1
	}

	dfs(root)
	return
}
