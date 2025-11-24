package main

import (
	"fmt"
	"sync/atomic"
)

/*
	每一个原子类型都会提供以下三个方法：
	Load()：原子的获取值
	Swap(newVal type) (old type)：原子的交换值，并且返回旧值
	Store(val type)：原子的存储值
*/

func main() {
	var aint64 atomic.Int64

	// 存储值
	aint64.Store(64)
	// 交换值
	var oldValue int64 = aint64.Swap(32)
	fmt.Println("old:", oldValue)
	// 增加
	aint64.Add(16)
	// 加载值
	fmt.Println(aint64.Load()) // 48

	/*
		也可以直接使用函数
			// 存储值
		   atomic.StoreInt64(&aint64, 64)
		   // 交换值
		   atomic.SwapInt64(&aint64, 128)
		   // 增加
		   atomic.AddInt64(&aint64, 112)
		   // 加载
		   fmt.Println(atomic.LoadInt64(&aint64))
	*/
}
