package offer专项突击_1_30

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, m int) *ListNode {
	tmp := head
	n := 0
	for tmp != nil {
		n++
		tmp = tmp.Next
	}
	k := n - m + 1
	pre := &ListNode{}
	t := pre
	pre.Next = head
	cnt := 0
	for pre != nil && cnt != k-1 {
		pre = pre.Next
		cnt++
	}
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	} else {
		pre.Next = nil
	}
	return t.Next
}
