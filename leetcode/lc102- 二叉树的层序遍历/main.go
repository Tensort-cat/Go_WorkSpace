package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	curLevel := []*TreeNode{root}

	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{}
		vals := []int{}
		for _, node := range curLevel {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		res = append(res, vals)
		curLevel = nextLevel
	}

	return res
}
