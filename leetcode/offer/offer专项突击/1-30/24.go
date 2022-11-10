package offer专项突击_1_30

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
