package offer专项突击_1_30

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	pre := 0
	var ans *ListNode
	for len(stack1) != 0 || len(stack2) != 0 || pre != 0 {
		topa, topb := 0, 0
		if len(stack1) != 0 {
			topa = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			topb = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		cur := topa + topb + pre
		pre = cur / 10
		cur = cur % 10
		// 使用头插法
		node := &ListNode{Val: cur}
		node.Next = ans
		ans = node
	}
	return ans
}

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
