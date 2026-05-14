package main

import (
	"fmt"
	"slices"
)

func main() {
	// 切片反转
	s1 := []int{1, 2, 3, 4, 5}
	slices.Reverse(s1)
	fmt.Println(s1)
}

/**
102. 二叉树的递归遍历
*/
// func levelOrder(root *TreeNode) [][]int {
// 	arr := [][]int{}

// 	depth := 0

// 	var order func(root *TreeNode, depth int)

// 	order = func(root *TreeNode, depth int) {
// 		if root == nil {
// 			return
// 		}
// 		if len(arr) == depth {
// 			arr = append(arr, []int{})
// 		}
// 		arr[depth] = append(arr[depth], root.Val)

// 		order(root.Left, depth+1)
// 		order(root.Right, depth+1)
// 	}

// 	order(root, depth)

// 	return arr
// }
