package main

import (
	"fmt"
	"reflect"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 通过反射可以获取函数的一切信息，也可以反射调用函数
func main() {
	/* 获取函数信息 */
	rType := reflect.TypeOf(Max)
	// 输出函数名称,字面量函数的类型没有名称
	fmt.Println(rType.Name())
	// 输出参数，返回值的数量
	fmt.Println(rType.NumIn(), rType.NumOut())
	rParamType := rType.In(0)
	// 输出第一个参数的类型
	fmt.Println(rParamType.Kind())
	rResType := rType.Out(0)
	// 输出第一个返回值的类型
	fmt.Println(rResType.Kind())

	/*
		通过反射值调用函数
		func (v Value) Call(in []Value) []Value
	*/
	// 传入参数数组
	rValue := reflect.ValueOf(Max)
	rResValue := rValue.Call([]reflect.Value{reflect.ValueOf(18), reflect.ValueOf(50)})
	for _, value := range rResValue {
		fmt.Println(value.Interface())
	}
}
