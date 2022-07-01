package _022_06

import (
	"math"
)

// https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
func findLUSlength(strs []string) int {
	maxl := -1
	for i := 0; i < len(strs); i++ {
		j := 0
		for ; j < len(strs); j++ {
			if j == i {
				continue
			}
			if isValid(strs[j], strs[i]) {
				break
			}
		}
		if j == len(strs) {
			maxl = int(math.Max(float64(maxl), float64(len(strs[i]))))
		}
	}
	return maxl
}

// t是否是s的子序列
func isValid(s, t string) bool {
	if len(s) < len(t) {
		return false
	}
	i, j := 0, 0
	for j != len(t) && i != len(s) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	if j == len(t) {
		return true
	}
	return false
}
