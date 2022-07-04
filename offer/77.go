package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// merge
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail := slow.Next
	slow.Next = nil
	l := sortList(head)
	r := sortList(tail)
	ans := mergeList(l, r)
	return ans
}

func mergeList(t1 *ListNode, t2 *ListNode) *ListNode {
	t3 := &ListNode{Val: -1}
	p := t3
	for t1 != nil && t2 != nil {
		if t1.Val < t2.Val {
			t3.Next = t1
			t3 = t3.Next
			t1 = t1.Next
		} else {
			t3.Next = t2
			t3 = t3.Next
			t2 = t2.Next
		}
	}
	if t1 != nil {
		t3.Next = t1
	}
	if t2 != nil {
		t3.Next = t2
	}
	return p.Next
}
