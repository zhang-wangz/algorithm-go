package offer

import "math"

// https://leetcode.cn/problems/wtcaE1/
func lengthOfLongestSubstring(s string) int {
	mp := map[byte]int{}
	left := 0
	right := 0
	ml := 0
	if s == "" {
		return ml
	}
	for right < len(s) {
		t := s[right]
		mp[t]++
		right++
		for mp[t] > 1 {
			ml = int(math.Max(float64(ml), float64(right-left-1)))
			t := s[left]
			mp[t]--
			left++
		}
	}
	if right == len(s) && mp[s[len(s)-1]] == 1 {
		return int(math.Max(float64(ml), float64(right-left)))
	}
	return ml
}
