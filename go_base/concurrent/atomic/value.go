package main

import (
	"fmt"
	"sync/atomic"
)

/*
atomic.Value 结构体，可以存储任意类型的值，结构体如下
	type Value struct {
		v any
	}
*/

func main() {
	var val atomic.Value
	val.Store(10)
	fmt.Println(val.Load())
	/*
			尽管可以存储任意类型，但是它不能存储 nil，并且前后存储的值类型应当一致，下面两个例子都无法通过编译
			example1:
			val.Store(nil)
			fmt.Println(val.Load()) // panic: sync/atomic: store of nil value into Value

			example2:
			val.Store("hello world")
		    val.Store(114154)
		    fmt.Println(val.Load()) // panic: sync/atomic: store of inconsistently typed value into Value
	*/
}
