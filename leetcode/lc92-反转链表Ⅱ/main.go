package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{Next: head}
	start, end := newHead, newHead
	var startPrev, endNext *ListNode
	for i := 0; i < left; i++ {
		startPrev = start
		start = start.Next
	}

	for i := 0; i < right; i++ {
		end = end.Next
		endNext = end.Next
	}

	// 断链
	startPrev.Next = nil
	end.Next = nil

	cur := start
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	startPrev.Next = prev
	start.Next = endNext

	return newHead.Next
}
