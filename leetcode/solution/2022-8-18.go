package main

// 最大相等频率
func maxEqualFreq(nums []int) (ans int) {
	maxFeq := 0
	feqs := map[int]int{}
	cnt := map[int]int{}
	for i, v := range nums {
		if cnt[v] > 0 {
			feqs[cnt[v]]--
		}
		cnt[v]++
		feqs[cnt[v]]++
		maxFeq = max(maxFeq, cnt[v])
		// 三种情况
		if maxFeq == 1 ||
			maxFeq*feqs[maxFeq]+(maxFeq-1)*feqs[maxFeq-1] == i+1 && feqs[maxFeq] == 1 ||
			maxFeq*feqs[maxFeq]+1 == i+1 && feqs[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
