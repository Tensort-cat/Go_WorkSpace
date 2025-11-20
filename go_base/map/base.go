package main

import "fmt"

func main() {
	// 创建方式1：使用make函数创建map  map_variable := make(map[KeyType]ValueType, initialCapacity)
	m := make(map[string]int) // 创建一个空的map
	fmt.Println(m)            // map[]

	m1 := make(map[string]int, 10) // 创建一个初始容量为10的map
	m1["apple"] = 100
	m1["banana"] = 200
	m1["orange"] = 300
	fmt.Println(m1) // map[]

	// 创建方式2：使用字面量语法创建map  map_variable := map[KeyType]ValueType{key1:value1, key2:value2, key3:value3}
	m2 := map[string]int{"apple": 100, "banana": 200, "orange": 300}
	fmt.Println(m2) // map[apple:100 banana:200 orange:300]

	// 获取键值对
	v1 := m1["apple"]
	fmt.Println(v1)      // 100
	v2, ok := m1["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	fmt.Println(v2, ok)  // 0 false

	// 获取map长度
	fmt.Println(len(m1)) // 3

	// 删除键值对
	delete(m1, "banana")
	fmt.Println(m1) // map[apple:100 orange:300]
}
