package swordReferoffer

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
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
	t = reverse(t)
	for pre != nil && t != nil {
		if pre.Val == t.Val {
			pre = pre.Next
			t = t.Next
		} else {
			return false
		}
	}
	return true
}
