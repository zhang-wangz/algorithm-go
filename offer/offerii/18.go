package main

// type ListNode struct {
//    Val int
//    Next *ListNode
//}

func deleteNode(head *ListNode, val int) *ListNode {
	tommy := &ListNode{}
	tommy.Next = head
	pre := tommy
	for head != nil {
		if head.Val == val {
			tommy.Next = head.Next
		}
		tommy = tommy.Next
		head = head.Next
	}
	return pre.Next
}
