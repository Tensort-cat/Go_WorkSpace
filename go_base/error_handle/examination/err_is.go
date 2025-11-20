// 任务：在 main 函数里用 errors.Is 判断 err 是否是 ErrPermission，并输出对应信息
package main

import (
	"errors"
	"fmt"
)

var ErrPermission = errors.New("无权限访问")

func checkAccess(user string) error {
	if user != "admin" {
		return fmt.Errorf("访问被拒绝: %w", ErrPermission)
	}
	return nil
}

func main() {
	err := checkAccess("guest")
	// TODO: 在这里判断是否是 ErrPermission
	// 如果是，打印 “权限错误”
	// 否则打印 “其他错误”
	if errors.Is(err, ErrPermission) {
		fmt.Println("权限错误")
	} else {
		fmt.Println("其他错误")
	}
}
