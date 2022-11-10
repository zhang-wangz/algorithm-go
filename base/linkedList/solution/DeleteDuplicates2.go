package solution

import "algorithm-go/base/linkedList"

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

func deleteDuplicates2(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return head
	}
	dummy := &linkedList.ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	var lastVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			// 存在[1,1,1]情况
			for head.Next != nil && head.Next.Val == lastVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
