package solution

import "algorithm-pattern/data_structure/linkedList"

// 反转从 m开始到n 节点的顺序
func ReverseList2(head *linkedList.ListNode, m, n int) *linkedList.ListNode {
	if head == nil {
		return head
	}
	pre := &linkedList.ListNode{Val: 0}
	pre.Next = head
	dummy := pre
	var i = 0
	for ; i < m-1; i++ {
		pre = pre.Next
		head = head.Next
	}
	// 此时 0->10(pre)->5(head)-> 7->8->9
	// head是开始要反转的地方
	// i 处于m-1位置
	// 假设m为2，n为4
	tmpPre := pre
	var p *linkedList.ListNode
	mid := head
	for ; i < n; i++ {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	// 此时 10->  7(pre)->5->nil  8(head)->9->nil
	tmpPre.Next = p
	mid.Next = head
	return dummy.Next
}
