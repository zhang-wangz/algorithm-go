package offer专项突击_1_30

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	pre := head
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	t := slow.Next
	slow.Next = nil
	tmp := reverse(t)

	for pre != nil && tmp != nil {
		l1tmp := pre.Next
		l2tmp := tmp.Next

		pre.Next = tmp
		pre = l1tmp

		tmp.Next = pre
		tmp = l2tmp
	}
	//node := &ListNode{}
	//flag := true
	//for pre != nil && tmp != nil {
	//	if flag {
	//		node.Next = pre
	//		pre = pre.Next
	//		node = node.Next
	//		flag = !flag
	//	} else {
	//		node.Next = tmp
	//		tmp = tmp.Next
	//		node = node.Next
	//		flag = !flag
	//	}
	//}
	//if pre != nil {
	//	node.Next = pre
	//}
	//if tmp != nil {
	//	node.Next = tmp
	//}
}
