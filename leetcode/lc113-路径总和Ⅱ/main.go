package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	sum := 0
	res := [][]int{}
	path := []int{}
	var ans func(*TreeNode)
	ans = func(root *TreeNode) {
		if root != nil {
			sum += root.Val
			path = append(path, root.Val)
			if root.Left == nil && root.Right == nil && sum == targetSum {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}

			ans(root.Left)
			ans(root.Right)
			sum -= root.Val
			path = path[:len(path)-1]
		}
	}

	ans(root)
	return res
}
