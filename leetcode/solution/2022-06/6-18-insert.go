package _022_06

import (
	"algorithm-go/linkedList"
	"math"
)

// 给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
//
//给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
//
//如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
//
//如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
func Insert(aNode *linkedList.ListNode, x int) *linkedList.ListNode {
	tmp := &linkedList.ListNode{Val: x}
	tmp.Next = tmp
	if aNode == nil {
		return tmp
	}
	ans := aNode
	min, max := aNode.Val, aNode.Val
	for aNode.Next != ans {
		aNode = aNode.Next
		v := aNode.Val
		max = int(math.Max(float64(max), float64(v)))
		min = int(math.Min(float64(min), float64(v)))
	}
	if max == min {
		tmp.Next = ans.Next
		ans.Next = tmp
		return ans
	} else {
		for !(aNode.Val == max && aNode.Next.Val == min) {
			aNode = aNode.Next
		}
		for !(x >= max || x <= min) && !(aNode.Val <= x && aNode.Next.Val >= x) {
			aNode = aNode.Next
		}
		tmp.Next = aNode.Next
		aNode.Next = tmp
	}
	return ans
}
