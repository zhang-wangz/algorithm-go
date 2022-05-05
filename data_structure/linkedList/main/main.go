package main

import (
	"algorithm-pattern/data_structure/linkedList"
	"algorithm-pattern/data_structure/linkedList/solution"
)

func main() {
	head := &linkedList.ListNode{Val: 10}
	head.Next = &linkedList.ListNode{Val: 5}
	head.Next.Next = &linkedList.ListNode{Val: 7}
	head.Next.Next.Next = &linkedList.ListNode{Val: 8}

	head = solution.ReverseList(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
