package offer

import "math"

//https://leetcode.cn/problems/M1oyTv/
func minWindow(s string, t string) string {
	need := map[byte]int{}
	win := map[byte]int{}
	if len(s) < len(t) {
		return ""
	}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	match := 0
	ans := [2]int{}
	minl := math.MaxInt
	left := 0
	for right := 0; right < len(s); right++ {
		t := s[right]
		if need[t] != 0 {
			win[t]++
			if win[t] == need[t] {
				match++
			}
		}
		for match == len(need) {
			if right-left+1 < minl {
				minl = right - left + 1
				ans[0] = left
				ans[1] = right
			}
			t := s[left]
			left++
			if need[t] != 0 {
				if win[t] == need[t] {
					match--
				}
				win[t]--
			}
		}
	}
	if minl == math.MaxInt {
		return ""
	}
	return s[ans[0] : ans[1]+1]
}
