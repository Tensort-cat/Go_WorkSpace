package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// signal.NotifyContext 会基于父 context 创建一个新的子 context。
	// 当进程收到指定的系统信号时，这个 ctx 就会被自动取消。
	//
	// 这里的 context.Background() 表示“根 context”，
	// 它本身不会主动结束，通常作为整个程序的起点。
	//
	// stop 是一个取消函数：
	// 1. 可以手动取消这个 ctx
	// 2. 释放 signal.NotifyContext 内部占用的资源
	//
	// 对初学者来说，可以把它理解成：
	// “把 Ctrl+C / 终止信号，转换成 <-ctx.Done() 可感知的退出事件”。
	//
	// 说明：
	// SIGINT 常见于 Ctrl+C
	// SIGTERM 常见于优雅终止
	// SIGKILL 在大多数系统中不能被程序捕获或处理，这里写上通常没有实际效果
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	// defer stop() 保证 main 退出前执行清理。
	// 即使程序是因为收到信号而结束，也建议这样写，形成固定习惯。
	defer stop()

	// 用一个无限循环模拟“服务持续运行”的场景。
	for {
		select {
		// ctx.Done() 会返回一个只读 channel。
		// 当 ctx 被取消时，这个 channel 会被关闭，
		// 所以 <-ctx.Done() 就会立刻收到通知。
		case <-ctx.Done():
			// 收到退出信号后，打印提示并结束程序。
			fmt.Println("terminate")
			return
		default:
			// default 表示“如果当前没有收到取消通知，就继续往下执行”。
			// 因为有 default，这个 select 不会阻塞。
		}

		// 模拟程序正在持续工作。
		fmt.Println("running")
		// 稍微休眠一下，避免空转导致 CPU 占用过高。
		time.Sleep(100 * time.Millisecond)
	}
}
