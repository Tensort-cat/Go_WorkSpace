package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(l *TreeNode, r *TreeNode) bool {
		if l == nil && r != nil {
			return false
		}
		if l != nil && r == nil {
			return false
		}
		if l == nil && r == nil {
			return true
		}
		if l.Val != r.Val {
			return false
		}

		return compare(l.Left, r.Right) && compare(l.Right, r.Left)
	}

	return compare(root.Left, root.Right)
}
