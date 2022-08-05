package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 最简单的当然是直接建立一个数组，然后随机
// 不用额外的空间，需要使用水塘抽样去计算，留在当前的概率
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (r *Solution) GetRandom() int {
	h := r.head
	sum, cur := 0, 1
	ans := -1
	for h != nil {
		// 随机0-1
		// 随机0-2
		// 随机0-3
		if rand.Intn(sum+1) < cur {
			ans = h.Val
		}
		sum += cur
		h = h.Next
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
