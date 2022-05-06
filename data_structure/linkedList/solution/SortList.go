package solution

import "algorithm-pattern/data_structure/linkedList"

// 给你链表的头结点head ，请将其按升序排列并返回排序后的链表
// 用mergeSort
func SortList(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMidNode(head)
	tail := mid.Next
	mid.Next = nil
	left := SortList(head)
	right := SortList(tail)
	res := MergeTwoLists(left, right)
	return res
}

func findMidNode(head *linkedList.ListNode) *linkedList.ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
