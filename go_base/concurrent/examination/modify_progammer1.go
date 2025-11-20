/*
以下程序存在并发问题，不一定能成功打印期待值(将0-9完整打印)：

	func main() {
		fmt.Println("start")
		for i := 0; i < 10; i++ {
			go fmt.Println(i)
		}
		fmt.Println("end")
	}

请修改它
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
	fmt.Println("end")
}
