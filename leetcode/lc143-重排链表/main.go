package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var reserve func(*ListNode) *ListNode
	reserve = func(head *ListNode) *ListNode {
		var prev *ListNode
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		return prev
	}

	cur := head.Next
	prev := head
	for cur != nil {
		prev.Next = nil
		prev.Next = reserve(cur)
		cur = prev.Next.Next
		prev = prev.Next
	}
}
