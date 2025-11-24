package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64 = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		// go Add1(1, &wg)
		go Add2(1, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func Add1(num int64, wg *sync.WaitGroup) { // 非原子操作 (不一定得到预期值)
	defer wg.Done()
	count += num
}

/*
对于 CAS 而言，有三个参数，内存值，期望值，新值。
执行时，CAS 会将期望值与当前内存值进行比较，如果内存值与期望值相同，就会执行后续的操作，
否则的话什么也不做。
对于 Go 中 atomic 包下的原子操作，CAS 相关的函数则需要传入地址，期望值，新值，
且会返回是否成功替换的布尔值。
*/
func Add2(num int64, wg *sync.WaitGroup) { // 使用乐观锁 (一定能得到预期值)
	defer wg.Done()
	for {
		expected := atomic.LoadInt64(&count)
		if atomic.CompareAndSwapInt64(&count, expected, expected+num) { // 尝试替换，若内存值与期望值相同，则替换为新值并返回 true，否则直接放弃返回false
			break
		}
	}
}
