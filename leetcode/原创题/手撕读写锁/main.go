package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
type RWLock struct {
	mu      sync.Mutex // 保护下面的变量
	readers int        // 正在读的读者数量
	writer  bool       // 是否有写者在写
	cond    *sync.Cond // 条件变量
}

func NewRWLock() *RWLock {
	rw := new(RWLock)
	rw.cond = sync.NewCond(&rw.mu)

	return rw
}

// 读锁
func (rw *RWLock) RLock() {
	rw.mu.Lock()
	for rw.writer { // 有写者在写
		rw.cond.Wait()
	}

	rw.readers++
	rw.mu.Unlock()
}

// 解读锁
func (rw *RWLock) RUnLock() {
	rw.mu.Lock()

	rw.readers--
	if rw.readers == 0 { // 自己是最后的读者
		rw.cond.Broadcast() // 唤醒写者
	}

	rw.mu.Unlock()
}

// 写锁
func (rw *RWLock) Lock() {
	rw.mu.Lock()

	for rw.readers > 0 || rw.writer { // 有读者或其他写者，阻塞
		rw.cond.Wait()
	}

	rw.writer = true // 防止其他写者或读者进入
	rw.mu.Unlock()
}

// 解写锁
func (rw *RWLock) Unlock() {
	rw.mu.Lock()

	rw.writer = false
	rw.cond.Broadcast()

	rw.mu.Unlock()
}

func main() {
	rw := NewRWLock()
	data := 0

	// 读协程
	for i := 0; i < 3; i++ {
		go func(id int) {
			time.Sleep(500 * time.Millisecond)

			rw.RLock()
			fmt.Printf("Reader %d: %d\n", id, data)
			rw.RUnLock()
		}(i)
	}

	// 写协程
	for i := 0; i < 2; i++ {
		go func(id int) {
			time.Sleep(500 * time.Millisecond)

			rw.Lock()
			data++
			fmt.Printf("Writer %d let data to be: %d\n", id, data)
			rw.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
}
