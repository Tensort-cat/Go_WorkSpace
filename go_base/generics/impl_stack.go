package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

// 入栈
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// 出栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	top := s.elements[len(s.elements)-1]
	return top, true
}

// 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	top1, ok := s.Pop()
	fmt.Println(top1, ok)
	top1, ok = s.Pop()
	fmt.Println(top1, ok)

	s2 := Stack[string]{}
	s2.Push("hello")
	s2.Push("world")
	top2, ok := s2.Pop()
	fmt.Println(top2, ok)
	top2, ok = s2.Pop()
	fmt.Println(top2, ok)
}
