/* func function_name( [parameter list] ) [return_types] {
   函数体
} */

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	a, b = swap(a, b)
	fmt.Printf("交换后 a = %d, b = %d\n", a, b)

}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y int) (int, int) {
	return y, x
}
