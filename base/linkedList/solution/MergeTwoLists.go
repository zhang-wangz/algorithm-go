package solution

import "algorithm-go/base/linkedList"

// 合成两个有序链表
func MergeTwoLists(l1 *linkedList.ListNode, l2 *linkedList.ListNode) *linkedList.ListNode {
	dummy := &linkedList.ListNode{Val: 0}
	res := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = l1
			dummy = dummy.Next
			l1 = l1.Next
		} else {
			dummy.Next = l2
			dummy = dummy.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		dummy.Next = l1
		dummy = dummy.Next
		l1 = l1.Next
	}
	for l2 != nil {
		dummy.Next = l2
		dummy = dummy.Next
		l2 = l2.Next
	}
	return res.Next
}
