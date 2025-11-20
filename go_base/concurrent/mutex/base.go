package main

import (
	"fmt"
	"sync"
)

// 全局变量模拟临界资源
var count int = 0

func adder(mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	count++
	mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	/*
		mutex.Lock()
		// 临界区start
		fmt.Println("Locked")
		// 临界区end
		mutex.Unlock()
	*/
	for i := 0; i < 10; i++ {
		go adder(&mutex, &wg)
	}
	wg.Wait()
	fmt.Println("count:", count)
}
