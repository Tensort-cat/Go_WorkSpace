// errors.Is 和 errors.As 用于处理错误链
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func getErr() error {
	return &MyError{
		Code: 404,
		Msg:  "page not found",
	}
}

func main() {
	err := findItem(1)
	// errors.Is 检查某个错误是否是特定错误或由该错误包装而成
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error:", err)
	}

	// errors.As 尝试将错误转换为特定类型
	var myErr *MyError
	// 以下会走else分支
	if errors.As(err, &myErr) {
		fmt.Println("MyError:", myErr.Code, myErr.Msg)
	} else {
		fmt.Println("Not a MyError:", err)
	}
}
