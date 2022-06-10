package solution

import (
	"algorithm-go/linkedList"
)

// https://leetcode.cn/problems/swap-nodes-in-pairs/
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *linkedList.ListNode) *linkedList.ListNode {
	var swap func(head *linkedList.ListNode) *linkedList.ListNode
	swap = func(head *linkedList.ListNode) *linkedList.ListNode {
		if head == nil || head.Next == nil {
			return nil
		}
		nextHead := head.Next.Next
		next := head.Next
		next.Next = head
		head.Next = swap(nextHead)
		return next
	}
	return swap(head)
}
