package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile() error {
	_, err := os.Open("notfound.txt")
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}
	return nil
}

func main() {
	err := readFile()
	// TODO: 判断 err 是否是 *os.PathError 类型
	// 如果是，打印：
	//   "文件路径错误:", pathErr.Path
	// 否则打印：
	//   "其他错误"
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("文件路径错误: ", pathErr.Path)
	} else {
		fmt.Println("其他错误")
	}
}
