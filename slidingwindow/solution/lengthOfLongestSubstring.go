package solution

import "math"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	ans := math.MinInt64
	win := map[byte]int{}
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		win[c]++

		for win[c] > 1 {
			c = s[left]
			left++
			win[c]--
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
