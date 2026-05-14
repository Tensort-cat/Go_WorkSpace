package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

/**
102. 二叉树的层序遍历：使用切片模拟队列，易理解
*/
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	curLevel := []*TreeNode{root} // 存放当前层节点
	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{} // 准备通过当前层生成下一层
		vals := []int{}

		for _, node := range curLevel {
			vals = append(vals, node.Val) // 收集当前层的值
			// 收集下一层的节点
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, vals)
		curLevel = nextLevel // 将下一层变成当前层
	}

	return
}

func main() {
	root := &TreeNode{
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
		Val:   5,
	}

	fmt.Println(levelOrder(root))
}
