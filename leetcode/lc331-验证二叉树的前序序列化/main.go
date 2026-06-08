package main

func isValidSerialization(preorder string) bool {
	// 从后往前遍历，#入栈，数字消耗两个#后生成一个#入栈
	stack := []byte{}
	for i := len(preorder) - 1; i >= 0; i-- {
		switch preorder[i] {
		case '#':
			stack = append(stack, '#')

		case ',':
			continue

		default:
			if len(stack) < 2 { // 栈中不存在两个#来抵消
				return false
			}
			stack = stack[:len(stack)-2] // 两个#出栈
			stack = append(stack, '#')   // 一个#入栈

			// i挪到非数字
			for i >= 0 && preorder[i] >= '0' && preorder[i] <= '9' {
				i--
			}
		}
	}

	return len(stack) == 1
}
