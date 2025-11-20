// sync.WaitGroup 用于等待多个 Goroutine 完成
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

/*
Add() 必须在启动 goroutine 之前调用
Done() 必须和每一个 Add(1) 一一对应, 因为WaitGroup 的内部计数器不能为负数, 直接触发panic
Wait() 必须在 Add() 和 Done() 都设置好后再调用
WaitGroup 不能重复使用，Wait之后计数器清零，不能再 Add
Wait() 相当于“封口”，计数器如果被改变，会立刻 panic
*/
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}
