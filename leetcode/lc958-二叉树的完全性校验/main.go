package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isCompleteTree(root *TreeNode) bool {
	level := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		level = append(level, node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	for i, node := range level {
		leftIndex := i*2 + 1
		rightIndex := i*2 + 2
		if leftIndex < len(level) && level[leftIndex] != node.Left {
			return false
		}
		if rightIndex < len(level) && level[rightIndex] != node.Right {
			return false
		}
	}

	return true
}
