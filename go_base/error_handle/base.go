package main

import (
	"errors"
	"fmt"
)

func main() {
	// 使用 errors 包创建错误
	err := errors.New("this is an error")
	fmt.Println(err) // 输出：this is an error

	/*
		在下面的例子中，我们在调用 Sqrt 的时候传递一个负数，
		然后就得到了 non-nil 的 error 对象，将此对象与 nil 比较，结果为 true，
		所以 fmt.Println(fmt 包在处理 error 时会调用 Error 方法)被调用，以输出错误
	*/
	result, err := Sqrt(-1)
	fmt.Println(result, err) // 输出：0 math: square root of negative number
	if err != nil {
		fmt.Println(err)
	}
	result, err = Sqrt(2)
	fmt.Println(result, err) // 输出：1.4142135623730951 <nil>
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 0, nil
}
