package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := len(nums)
	L := make([]int, l, l)
	R := make([]int, l, l)
	L[0], R[l-1] = 1, 1

	for i := 1; i < l; i++ {
		L[i] = nums[i-1] * L[i-1]
		R[l-i-1] = nums[l-i] * R[l-i]
	}
	fmt.Println(L, R)

	var answer []int
	for i := 0; i < l; i++ {
		answer = append(answer, L[i]*R[i])
	}
	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 0}))
}
