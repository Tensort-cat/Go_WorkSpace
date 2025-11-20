package main

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64      // 面积
	Perimeter() float64 // 周长
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	c := Circle{Radius: 5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	/* 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
	任意类型都实现了空接口。
	常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。 */
	printValue(42)          // int
	printValue("hello")     // string
	printValue(3.14)        // float64
	printValue([]int{1, 2}) // slice
	printValue(s)

	// 类型断言
	var i interface{} = "Hello"
	str, ok := i.(string) // 类型断言，ok 表示断言是否成功
	if ok {
		fmt.Println(str)
	}
}

func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}
