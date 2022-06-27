package swordReferoffer

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
	node := &ListNode{}
	res := node
	flag := true
	for pre != nil && tmp != nil {
		if flag {
			node.Next = pre
			pre = pre.Next
			node = node.Next
			flag = !flag
		} else {
			node.Next = tmp
			tmp = tmp.Next
			node = node.Next
			flag = !flag
		}
	}
	if pre != nil {
		node.Next = pre
		pre = pre.Next
		node = node.Next
	}

	if tmp != nil {
		node.Next = tmp
		tmp = tmp.Next
		node = node.Next
	}
	head = res.Next
}
