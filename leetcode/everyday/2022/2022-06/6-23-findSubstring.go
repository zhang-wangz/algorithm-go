package _022_06

// FindSubstring https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
//
//	给定一个字符串s和一些长度相同的单词words 。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
//
// 注意子串要与words中的单词完全匹配，中间不能有其他字符 ，但不需要考虑words中单词串联的顺序。
func FindSubstring(s string, words []string) []int {
	mp := map[string]int{}
	for _, v := range words {
		mp[v] += 1
	}
	le := len(words[0])
	lewords := le * len(words)
	if len(s) < lewords {
		return []int{}
	}
	ans := make([]int, 0)
	for i := 0; i <= len(s)-lewords; i++ {
		// 定长窗口，n*le
		win := map[string]int{}
		err := false
		for j := 0; j < len(words); j++ {
			cur := s[i+j*le : i+j*le+le]
			if _, ok := mp[cur]; !ok {
				err = true
				break
			}
			win[cur]++
			if win[cur] > mp[cur] {
				err = true
				break
			}
		}
		if !err {
			ans = append(ans, i)
		}
	}
	return ans
}
