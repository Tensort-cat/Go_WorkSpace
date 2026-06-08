package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一，先统计长度，找到倒数第N个点再删除
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	l := 0
	p := head
	for p != nil {
		l++
		p = p.Next
	}
	if l == 1 {
		return nil
	}

	var prev *ListNode
	cur := head
	for i := 0; i < l-n; i++ {
		prev = cur
		cur = cur.Next
	}

	if cur == head {
		return head.Next
	}
	prev.Next = cur.Next
	return head
}

// 方法二：快慢指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, quick := dummy, dummy

	for i := 0; i < n; i++ {
		quick = quick.Next
	}

	var prev *ListNode
	for quick != nil {
		prev = slow
		slow = slow.Next
		quick = quick.Next
	}

	if slow == head {
		return head.Next
	}
	prev.Next = slow.Next
	return head
}
