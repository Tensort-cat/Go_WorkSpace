/*
互斥锁适合读操作与写操作频率都差不多的情况，对于一些读多写少的数据，如果使用互斥锁，会造成大量的不必要的协程竞争锁，这会消耗很多的系统资源，这时候就需要用到读写锁，即读写互斥锁，对于一个协程而言：

	如果获得了读锁，其他协程进行写操作时会阻塞，其他协程进行读操作时不会阻塞
	如果获得了写锁，其他协程进行写操作时会阻塞，其他协程进行读操作时会阻塞

// 加读锁
func (rw *RWMutex) RLock()

// 尝试加读锁
func (rw *RWMutex) TryRLock() bool

// 解读锁
func (rw *RWMutex) RUnlock()

// 加写锁
func (rw *RWMutex) Lock()

// 尝试加写锁
func (rw *RWMutex) TryLock() bool

// 解写锁
func (rw *RWMutex) Unlock()
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count = 0

var rw sync.RWMutex

func main() {
	wait.Add(12)
	// 读多写少
	go func() {
		for i := 0; i < 3; i++ {
			go Write(&count)
		}
		wait.Done()
	}()
	go func() {
		for i := 0; i < 7; i++ {
			go Read(&count)
		}
		wait.Done()
	}()
	// 等待子协程结束
	wait.Wait()
	fmt.Println("最终结果", count)
}

func Read(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	rw.RLock()
	fmt.Println("拿到读锁")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println("释放读锁", *i)
	rw.RUnlock()
	wait.Done()
}

func Write(i *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	rw.Lock()
	fmt.Println("拿到写锁")
	temp := *i
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	*i = temp + 1
	fmt.Println("释放写锁", *i)
	rw.Unlock()
	wait.Done()
}
