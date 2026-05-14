package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sort.Ints(nums) // 升序排序
	fmt.Printf("升序排序: %v\n", nums)

	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // 降序排序
	fmt.Printf("降序排序: %v\n", nums)

	// 自定义排序
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age // 按年龄升序排序
	})
	fmt.Printf("按年龄升序排序: %v\n", people)
}
