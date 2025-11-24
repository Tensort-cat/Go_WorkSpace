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

// 条件变量
var cond = sync.NewCond(rw.RLocker())

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
	// 条件不满足就一直阻塞
	for *i < 3 {
		cond.Wait()
	}
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
	// 唤醒所有因条件变量阻塞的协程
	cond.Broadcast()
	wait.Done()
}
