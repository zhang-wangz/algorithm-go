package solution

// https://leetcode.cn/problems/minimum-window-substring/
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

func minWindow(s string, t string) string {
	win := map[byte]int{}
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	min := 1<<63 - 1
	left, right := 0, 0
	start, end := 0, 0
	match := 0
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				start = left
				end = right
				min = right - left
			}
			c = s[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == 1<<61-1 {
		return ""
	}
	return s[start : end+1]
}
