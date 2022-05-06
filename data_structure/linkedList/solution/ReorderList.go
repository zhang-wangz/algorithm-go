package solution

import (
	"algorithm-pattern/data_structure/linkedList"
)

// https://leetcode-cn.com/problems/reorder-list/
// 给定一个单链表 L：L→L→…→L__n→L 将其重新排列后变为： L→L__n→L→L__n→L→L__n→…
func ReorderList(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return nil
	}
	mid := findMidNode(head)
	tail := ReverseList(mid.Next)
	mid.Next = nil
	head = mergeTwoList(head, tail)
	return head
}

func mergeTwoList(l1, l2 *linkedList.ListNode) *linkedList.ListNode {
	dummy := &linkedList.ListNode{Val: 0}
	res := dummy
	flag := true
	for l1 != nil && l2 != nil {
		if flag {
			dummy.Next = l1
			dummy = dummy.Next
			l1 = l1.Next
			flag = !flag
		} else {
			dummy.Next = l2
			dummy = dummy.Next
			l2 = l2.Next
			flag = !flag
		}
	}
	if l1 != nil {
		dummy.Next = l1
		dummy = dummy.Next
		l1 = l1.Next
	}
	if l2 != nil {
		dummy.Next = l2
		dummy = dummy.Next
		l2 = l2.Next
	}
	return res.Next
}
