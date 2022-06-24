package swordReferoffer

// https://leetcode.cn/problems/MPnaiL/
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
func checkInclusion(pattern string, str string) bool {
	need := map[byte]int{}
	for i := 0; i < len(pattern); i++ {
		need[pattern[i]]++
	}
	for i := 0; i <= len(str)-len(pattern); i++ {
		err := true
		win := map[byte]int{}
		for j := 0; j < len(pattern); j++ {
			t := str[i+j]
			if need[t] == 0 {
				err = false
				break
			}
			win[t]++
			if win[t] > need[t] {
				err = false
				break
			}
		}
		if err {
			return true
		}
	}
	return true
}
