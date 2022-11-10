package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 弄成一个环
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	it := head
	for it.Next != nil {
		it = it.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	it.Next = head
	for add > 0 {
		it = it.Next
		add--
	}
	res := it.Next
	it.Next = nil
	return res
}
