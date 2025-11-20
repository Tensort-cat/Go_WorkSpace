package main

import (
	"fmt"
)

type DivideError struct {
	Dividend int // 被除数
	Divisor  int // 除数
}

/*
error 接口定义如下：
type error interface {
	Error() string
}
*/
// 实现 error 接口的 Error 方法
func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err) // 输出：cannot divide 10 by 0
	}
}
