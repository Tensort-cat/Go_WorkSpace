package main

import (
	"fmt"
	"slices"
)

func decodeString(s string) string {
	// 数字：更新num
	// 左括号，入栈
	// 右括号，出栈到左括号，将单词重复num次
	// 字母：栈为空时直接输出，非空入栈
	stack := []byte{}
	numStack := []int{}
	byteRes := []byte{}
	for i := 0; i < len(s); {
		if s[i] >= '1' && s[i] <= '9' {
			num := 0
			j := i
			for ; s[j] >= '0' && s[j] <= '9'; j++ {
				num = num*10 + int(s[j]-'0')
			}
			numStack = append(numStack, num)
			i = j
			continue
		} else if s[i] == '[' {
			stack = append(stack, s[i])
		} else if s[i] == ']' {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			word := []byte{}
			for b != '[' {
				word = append(word, b)
				b = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			slices.Reverse(word)
			if len(stack) == 0 {
				for j := 0; j < num; j++ {
					byteRes = append(byteRes, word...)
				}
			} else {
				for j := 0; j < num; j++ {
					stack = append(stack, word...)
				}
			}

		} else { // 字母
			if len(stack) == 0 {
				byteRes = append(byteRes, s[i])
			} else {
				stack = append(stack, s[i])
			}
		}
		i++
	}

	return string(byteRes)
}

func main() {
	s := "11[fuck]"
	fmt.Printf("decodeString(s): %v\n", decodeString(s))
}
