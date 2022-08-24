package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 前缀和求法，当出现两次相同的sum时候，说明这中间的值为0，直接连接即可
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	m := map[int]*ListNode{}
	sum := 0
	for d := dummy; d != nil; d = d.Next {
		sum += d.Val
		m[sum] = d
	}
	sum = 0
	for d := dummy; d != nil; d = d.Next {
		sum += d.Val
		d.Next = m[sum].Next
	}
	return dummy.Next
}
