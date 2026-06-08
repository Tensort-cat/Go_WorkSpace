package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	prev, cur := newHead, head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur.Next = cur.Next.Next
			}
			// 删自己
			cur = cur.Next
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
		}

	}

	return newHead.Next
}
