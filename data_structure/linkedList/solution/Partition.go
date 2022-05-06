package solution

import "algorithm-pattern/data_structure/linkedList"

// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
// https://leetcode-cn.com/problems/partition-list/
// 我想的是保存最大结点，小的结点放到前面，官方题解是把大结点保存到另一个链表，最后进行拼接
func Partition(head *linkedList.ListNode, x int) *linkedList.ListNode {
	if head == nil {
		return head
	}
	dummy := &linkedList.ListNode{Val: -200}
	dummy.Next = head
	// pre, head, pre在head之前
	// x = 4
	// 0(preMax)->4(headMax)->3(pre)->2(head)->5
	pre := dummy
	preMax := dummy
	headMax := head
	// 标记第一个大于x的值
	flag := 1
	for head != nil {
		if head.Val >= x && flag == 1 {
			preMax = pre
			headMax = head
			flag = 0
			head = head.Next
			pre = pre.Next
		} else if head.Val >= x && flag == 0 {
			head = head.Next
			pre = pre.Next
		} else if head.Val < x && flag == 0 {
			tmp := head.Next
			pre.Next = head.Next
			preMax.Next = head
			// 因为结点到前面去了，所以导致大的结点前面的结点也变了，需要改变preMax
			preMax = head
			head.Next = headMax
			head = tmp
		} else if head.Val < x && flag == 1 {
			head = head.Next
			pre = pre.Next
		}
	}
	return dummy.Next
}
