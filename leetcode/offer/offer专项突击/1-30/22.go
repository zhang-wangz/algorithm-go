package offer专项突击_1_30

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	t := head
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == fast {
		slow = slow.Next
		fast = t
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}
