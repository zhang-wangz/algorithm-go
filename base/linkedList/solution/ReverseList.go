package solution

import "algorithm-go/base/linkedList"

func ReverseList(head *linkedList.ListNode) *linkedList.ListNode {
	var prev *linkedList.ListNode
	if head == nil {
		return head
	}
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
