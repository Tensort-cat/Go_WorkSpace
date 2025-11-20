package main

import "fmt"

func main() {
	ch := make(chan int, 3) // 带缓冲通道

	// 生产者
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("生产者发送：", i)
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	for v := range ch {
		fmt.Println("消费者接收：", v)
	}

	fmt.Println("所有数据处理完毕！")
}
