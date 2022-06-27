package swordReferoffer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := map[*ListNode]bool{}
	for headA != nil && headB != nil {
		if !mp[headA] {
			mp[headA] = true
		} else {
			return headA
		}
		if !mp[headB] {
			mp[headB] = true
		} else {
			return headB
		}
		headA = headA.Next
		headB = headB.Next
	}
	for headA != nil {
		if !mp[headA] {
			mp[headA] = true
		} else {
			return headA
		}
		headA = headA.Next
	}
	for headB != nil {
		if !mp[headB] {
			mp[headB] = true
		} else {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
