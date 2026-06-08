package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var ans func(*TreeNode, int, int) bool
	ans = func(root *TreeNode, lower int, upper int) bool {
		if root == nil {
			return true
		}

		if root.Val <= lower || root.Val >= upper {
			return false
		}

		return ans(root.Left, lower, root.Val) && ans(root.Right, root.Val, upper)
	}

	return ans(root, math.MinInt, math.MaxInt)
}
