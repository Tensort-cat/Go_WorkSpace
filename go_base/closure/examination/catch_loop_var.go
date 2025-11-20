// 下面的代码有什么问题？请修改，使它正确打印 0, 1, 2
package main

import "fmt"

func main() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		j := i
		funcs = append(funcs, func() {
			fmt.Println(j)
		})
	}

	for _, f := range funcs {
		f() // 期望输出 0, 1, 2
	}
}
