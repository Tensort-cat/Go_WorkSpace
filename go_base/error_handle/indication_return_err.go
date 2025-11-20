// Go 中，错误通常作为函数的返回值返回，开发者需要显式检查并处理
package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 10, 0
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
