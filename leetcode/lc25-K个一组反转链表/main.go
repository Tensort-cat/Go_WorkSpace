package main

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

func reverseKGroup(head *ListNode, k int) *ListNode {
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

func main() {

}
