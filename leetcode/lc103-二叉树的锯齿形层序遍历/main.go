package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	isPre := true
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

		if !isPre {
			slices.Reverse(vals)
		}
		res = append(res, vals)
		curLevel = nextLevel
		isPre = !isPre
	}

	return res
}
