package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	heads := []*ListNode{}
	p := head
	var prev *ListNode
	for p != nil {
		heads = append(heads, p)
		for i := 0; i < k && p != nil; i++ {
			prev = p
			p = p.Next
		}
		// 断链
		prev.Next = nil
	}

	newHeads := []*ListNode{}
	for i := 0; i < len(heads)-1; i++ {
		newHeads = append(newHeads, reverse(heads[i]))
	}
	// 看看最后一组需不需要反转
	l := 0
	p = heads[len(heads)-1]
	for p != nil {
		p = p.Next
		l++
	}
	if l == k {
		newHeads = append(newHeads, reverse(heads[len(heads)-1]))
	} else {
		newHeads = append(newHeads, heads[len(heads)-1])
	}

	// heads变成了新的每个子链表尾指针
	for i := 0; i < len(heads)-1; i++ {
		heads[i].Next = newHeads[i+1]
	}

	return newHeads[0]
}

func travel(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	cnt := 0
	p := head
	var q *ListNode
	for p != nil && cnt < k {
		q = p
		p = p.Next
		cnt++
	}
	if cnt < k {
		return head
	}

	// 反转
	q.Next = nil // 断链
	nextHead := reverseKGroup2(p, k)
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head.Next = nextHead
	return prev
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	head := new(ListNode)
	p := head
	var val int
	for i := 0; i < n-1; i++ {
		fmt.Scan(&val)
		p.Val = val
		p.Next = new(ListNode)
		p = p.Next
	}
	fmt.Scan(&val)
	p.Val = val
	p.Next = nil

	travel(head)
	// newHead := reverseKGroup1(head, k)
	newHead := reverseKGroup2(head, k)

	travel(newHead)
}
