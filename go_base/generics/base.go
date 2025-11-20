/*
// 基本语法结构

	func 函数名[T 约束](参数 T) 返回值类型 {
	    // 函数体
	}

	type 类型名[T 约束] struct {
	    // 结构体字段
	}
*/
package main

import "fmt"

func main() {
	// any约束
	PrintAny(12)
	PrintAny("你好")
	PrintAny(3.1415926)

	// comparable约束
	arrInt := []int{1, 2, 3, 4, 5}
	fmt.Println(FindIndex(arrInt, 3))
	arrStr := []string{"apple", "banana", "orange"}
	fmt.Println(FindIndex(arrStr, "banana"))

	// 联合约束
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.1, 2.2))

	// 自定义方法约束
	wangcai := Dog{"WangCai"}
	AnimalSpeak(wangcai)
}

// go原生提供约束
// 1. any约束: any 是空接口 interface{} 的别名，表示任何类型都可以
func PrintAny[T any](v T) {
	fmt.Printf("value: %v, type: %T\n", v, v)
}

// 2. comparable约束: 表示类型可以进行比较运算
func FindIndex[T comparable](arr []T, val T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

// 3. 联合约束：声明接口，使用 | 运算符组合多个类型
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// 自定义约束
// 1. 方法约束
type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func AnimalSpeak[T Animal](animal T) {
	fmt.Println(animal.Speak())
}
