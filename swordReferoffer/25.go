package swordReferoffer

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	t := ans
	stack1 := make([]*ListNode, 0)
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next
	}
	stack2 := make([]*ListNode, 0)
	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	pre := 0
	for len(stack1) != 0 && len(stack2) != 0 {
		v := stack1[len(stack1)-1].Val + stack2[len(stack2)-1].Val + pre
		pre = v / 10
		v = v % 10
		ans.Next = &ListNode{Val: v}
		ans = ans.Next
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
	}
	for len(stack1) != 0 {
		v := stack1[len(stack1)-1].Val + pre
		pre = v / 10
		v = v % 10
		ans.Next = &ListNode{Val: v}
		ans = ans.Next
		stack1 = stack1[:len(stack1)-1]
	}

	for len(stack2) != 0 {
		v := stack2[len(stack2)-1].Val + pre
		pre = v / 10
		v = v % 10
		ans.Next = &ListNode{Val: v}
		ans = ans.Next
		stack2 = stack2[:len(stack2)-1]
	}
	if pre != 0 {
		ans.Next = &ListNode{Val: pre}
		ans = ans.Next
	}
	return reverse(t.Next)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		t := head.Next
		head.Next = pre
		pre = head
		head = t
	}
	return pre
}
