package offer

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		t := head.Next
		head.Next = pre
		pre = head
		head = t
	}
	return pre
}
