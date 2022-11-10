package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
	}
	for i, j := 0, len(ans)-1; i != j; i++ {
		ans[i], ans[j] = ans[j], ans[i]
		j--
	}
	return ans
}
