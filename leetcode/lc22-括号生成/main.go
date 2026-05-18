package main

func generateParenthesis(n int) []string {
	var backtracking func(start int)
	res := []string{}
	var brackets string
	backtracking = func(start int) {
		if len(brackets) == n*2 {
			if isVaild(brackets) {
				tmp := brackets
				res = append(res, tmp)
			}
			return
		}
		for i := start; i < n*2; i++ {
			brackets += "("
			backtracking(i + 1)
			brackets = brackets[:len(brackets)-1]
			brackets += ")"
			backtracking(i + 1)
			brackets = brackets[:len(brackets)-1]
		}
	}

	backtracking(0)
	return res
}

func isVaild(brackets string) bool {
	stack := []byte{}
	for i := 0; i < len(brackets); i++ {
		if brackets[i] == '(' {
			stack = append(stack, brackets[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
