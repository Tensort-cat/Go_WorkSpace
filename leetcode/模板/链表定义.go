package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func initLinkList() *ListNode {
	return &ListNode{Next: nil}
}

func (l *ListNode) travel() {
	p := l.Next
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
}

func main() {
	head := initLinkList()
	tail := head
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)
		tail.Next = &ListNode{
			Next: nil,
			Val:  val,
		}
		tail = tail.Next
	}

	head.travel()
}
