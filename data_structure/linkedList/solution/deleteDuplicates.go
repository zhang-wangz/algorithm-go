package solution

import "algorithm-pattern/data_structure/linkedList"

func deleteDuplicates(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return head
	}
	node := head
	for head != nil {
		// 存在[1,1,1]这种情况
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return node
}
