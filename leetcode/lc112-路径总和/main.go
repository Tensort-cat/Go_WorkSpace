package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	res := false
	var ans func(*TreeNode)
	ans = func(root *TreeNode) {
		if root != nil {
			sum += root.Val
			if root.Left == nil && root.Right == nil && sum == targetSum {
				res = true
				return
			}

			ans(root.Left)
			ans(root.Right)
			sum -= root.Val
		}
	}

	ans(root)
	return res
}
