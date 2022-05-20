package sequence

// https://leetcode.cn/problems/word-break/
// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
// 判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		m[v] = true
	}
	dp := make([]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = false
	}
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if dp[i] {
			for j := i + 1; j <= len(s); j++ {
				if ok, _ := m[s[i:j]]; ok {
					dp[j] = true
				}
			}
		}
	}
	return dp[len(s)]
}
