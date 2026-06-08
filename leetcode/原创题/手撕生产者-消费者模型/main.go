package main

import (
	"fmt"
	"time"
)

var ch chan int

// 生产者
func Producer(id int) {
	for {
		ch <- 1
	}

	close(ch)
}

// 消费者
func Consumer(id int) {
	for msg := range ch {
		fmt.Printf("Consumer %d: %d\n", id, msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch = make(chan int, 5)

	go Producer(114514)

	for i := 0; i < 3; i++ {
		go Consumer(i)
	}

	for {
	}
}
