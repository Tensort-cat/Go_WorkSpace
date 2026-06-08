package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	// 空格：栈空或栈顶为空格抛弃，栈非空且栈顶非空格入栈
	// 字母：直接入栈
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			if len(stack) == 0 || stack[len(stack)-1] == ' ' {
				continue
			}
			stack = append(stack, ' ')

		default:
			stack = append(stack, s[i])
		}
	}

	// 最后末尾可能有空格
	if stack[len(stack)-1] == ' ' {
		stack = stack[:len(stack)-1]
	}

	res := strings.Split(string(stack), " ")
	slices.Reverse(res)
	return strings.Join(res, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("reverseWords(s): %v\n", reverseWords(s))
}
