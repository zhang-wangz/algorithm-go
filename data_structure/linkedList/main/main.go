package main

import (
	"algorithm-pattern/data_structure/linkedList"
	"algorithm-pattern/data_structure/linkedList/solution"
)

func main() {
	head := &linkedList.ListNode{Val: 1}
	head.Next = &linkedList.ListNode{Val: 4}
	head.Next.Next = &linkedList.ListNode{Val: 3}
	head.Next.Next.Next = &linkedList.ListNode{Val: 0}
	head.Next.Next.Next.Next = &linkedList.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &linkedList.ListNode{Val: 2}

	head = solution.SortList(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
