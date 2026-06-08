package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var head *TreeNode

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 右 - 左 - 根
		dfs(node.Right)
		dfs(node.Left)
		node.Left = nil
		node.Right = head // 头插法，相当于链表的 node.Next = head
		head = node       // 现在链表头节点是 node
	}

	dfs(root)
}
