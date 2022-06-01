package solution

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

func findAnagrams(s string, p string) []int {
	win := map[byte]int{}
	need := map[byte]int{}
	res := make([]int, 0)
	left, right := 0, 0
	match := 0
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
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

		for match == len(need) && right-left >= len(p) {
			if right-left == len(p) && match == len(need) {
				res = append(res, left)
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
	return res
}
