package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return count(root, targetSum) +
		pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum)
}

func count(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}

	res := 0

	if root.Val == target {
		res++
	}

	res += count(root.Left, target-root.Val)
	res += count(root.Right, target-root.Val)

	return res
}
