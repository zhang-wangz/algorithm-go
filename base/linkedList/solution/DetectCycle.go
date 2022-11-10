package solution

import "algorithm-go/base/linkedList"

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
func detectCycle(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		// 如果相遇回到起点
		// slow得走一格
		// b+a+f-1+a = 2f+2a => b-1=f
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return nil
}
