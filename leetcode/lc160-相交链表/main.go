package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != nil {
		for q != nil {
			if p == q {
				return p
			}
			q = q.Next
		}
		p = p.Next
		q = headB
	}
	return nil
}

// 我们必能相遇
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}

	return p
}
