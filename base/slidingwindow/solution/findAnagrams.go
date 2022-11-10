package solution

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 给定两个字符串s和p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

func findAnagrams(s string, p string) (res []int) {
	win := map[byte]int{}
	need := map[byte]int{}
	match := 0
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	l := 0
	for r := range s {
		c := s[r]
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		for match == len(need) && r-l+1 >= len(p) {
			if r-l+1 == len(p) && match == len(need) {
				res = append(res, l)
			}
			c = s[l]
			l++
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
