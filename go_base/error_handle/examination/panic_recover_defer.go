// 模拟打开文件失败但程序不崩溃
package main

import "fmt"

func openFile() {
	fmt.Println("尝试打开文件...")
	panic("文件不存在！") // 模拟严重错误
}

func safeOpen() {
	// TODO: 请在这里添加 defer + recover
	// 要求：捕获 panic 并打印 “捕获到错误: …”
	// 然后程序能继续执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到错误:", r)
		}
	}()
	openFile()
}

func main() {
	fmt.Println("程序开始")
	safeOpen()
	fmt.Println("程序结束后仍能正常运行")
}
