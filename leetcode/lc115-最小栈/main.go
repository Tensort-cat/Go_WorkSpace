package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {

	this.stack = append(this.stack, val)

	// minStack 为空
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
		return
	}

	// 当前最小值
	curMin := this.minStack[len(this.minStack)-1]

	this.minStack = append(this.minStack, min(curMin, val))
}

func (this *MinStack) Pop() {

	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
