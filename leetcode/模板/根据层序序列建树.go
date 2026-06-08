package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入格式示例：
// 1 2 3 nil nil 4 5
// 表示：
//       1
//     /   \
//    2     3
//         / \
//        4   5

func build(level []string, idx int) *TreeNode {
	// 越界
	if idx >= len(level) {
		return nil
	}

	// 空节点
	if level[idx] == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(level[idx])
	root := &TreeNode{
		Val: val,
	}

	// 递归建立左右子树
	root.Left = build(level, idx*2+1)
	root.Right = build(level, idx*2+2)

	return root
}

// 前序遍历验证
func preorder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	level := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&level[i])
	}

	root := build(level, 0)

	preorder(root)
}
