package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	l := make(chan struct{})

	go send(ch1)
	go send(ch2)
	go send(ch3)

	// time.After 函数，其返回值是一个只读的管道，该函数配合 select 使用可以非常简单的实现超时机制
	go func() {
	Loop:
		for {
			select {
			case n, ok := <-ch1:
				fmt.Println("ch1", n, ok)
			case n, ok := <-ch2:
				fmt.Println("ch2", n, ok)
			case n, ok := <-ch3:
				fmt.Println("ch3", n, ok)
			case <-time.After(time.Second): // 超时退出
				break Loop
			}
		}
		// 让主线程退出
		l <- struct{}{}
	}()
	// 等待开启的goroutine结束
	<-l
}

func send(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- i
	}
}
