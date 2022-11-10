package main

import (
	"algorithm-go/base/linkedList"
	"algorithm-go/base/linkedList/solution"
)

func main() {
	head := &linkedList.ListNode{Val: 1}
	head.Next = &linkedList.ListNode{Val: 4}
	head.Next.Next = &linkedList.ListNode{Val: 3}
	head.Next.Next.Next = &linkedList.ListNode{Val: 0}
	head.Next.Next.Next.Next = &linkedList.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &linkedList.ListNode{Val: 2}

	pre := solution.ReorderList(head)
	for pre != nil {
		println(pre.Val)
		pre = pre.Next
	}
}
