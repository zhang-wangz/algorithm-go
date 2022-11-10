package main

// https: //leetcode.cn/problems/shortest-impossible-sequence-of-rolls/
// 就是求互不交叉的完整序列出现了几次
func shortestSequence(rolls []int, k int) int {
	set := map[int]struct{}{}
	ans := 0
	for i := 0; i < len(rolls); i++ {
		set[rolls[i]] = struct{}{}
		if len(set) == k {
			ans++
			set = map[int]struct{}{}
		}
	}
	return ans + 1
}
