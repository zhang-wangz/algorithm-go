package swordReferoffer

// https://leetcode.cn/problems/MPnaiL/
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
// 维护固定长度
func checkInclusion(pattern string, str string) bool {
	need := [26]int{}
	win := [26]int{}
	if len(pattern) > len(str) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		need[pattern[i]-'a']++
		win[str[i]-'a']++
	}
	if need == win {
		return true
	}
	// 维护固定长度len(pattern)
	for i := len(pattern); i < len(str); i++ {
		win[str[i]-'a']++
		win[str[i-len(pattern)]-'a']--
		if need == win {
			return true
		}
	}
	return false
}

// 变长窗口求解
func checkInclusion2(pattern string, str string) bool {
	if len(pattern) > len(str) {
		return false
	}
	win := [26]int{}
	for i := 0; i < len(pattern); i++ {
		win[pattern[i]-'a']--
	}
	left := 0
	for right := 0; right < len(str); right++ {
		t := str[right] - 'a'
		win[t]++
		for win[t] > 0 {
			win[str[left]-'a']--
			left++
		}
		if right-left+1 == len(pattern) {
			return true
		}
	}
	return false
}
