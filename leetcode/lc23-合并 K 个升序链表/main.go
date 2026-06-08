package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) != 1 {
		tmp := []*ListNode{}
		for i := 0; i < len(lists)-1; i += 2 {
			tmp = append(tmp, merge(lists[i], lists[i+1]))
		}
		if len(lists)%2 != 0 {
			tmp = append(tmp, lists[len(lists)-1])
		}

		lists = tmp
	}

	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}

	l2.Next = merge(l1, l2.Next)
	return l2
}
