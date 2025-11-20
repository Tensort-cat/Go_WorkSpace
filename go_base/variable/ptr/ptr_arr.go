// 指针数组
package main

import "fmt"

const MAX_SIZE int = 3

func main() {

	arr := [MAX_SIZE]int{1, 2, 3}
	for i := 0; i < MAX_SIZE; i++ {
		fmt.Println(arr[i])
	}

	// ptr 为整型指针数组。因此每个元素都指向了一个值
	var ptr [MAX_SIZE]*int
	for i := 0; i < MAX_SIZE; i++ {
		ptr[i] = &arr[i] /* 整数地址赋值给指针数组 */
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
