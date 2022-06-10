package solution

import "algorithm-go/linkedList"

// 是不是回文链表
func isPalindrome(head *linkedList.ListNode) bool {
	if head == nil {
		return true
	}
	p := head
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 此时slow是中点，只是验证需要使用中点的下一个
	// 比如 1221, 此时head需要指向2
	tailHead := slow.Next
	tail := reserve(tailHead)
	slow.Next = nil
	for p != nil && tail != nil {
		if p.Val != tail.Val {
			return false
		}
		p = p.Next
		tail = tail.Next
	}
	return true
}

func reserve(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return head
	}
	var p *linkedList.ListNode
	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p
}
