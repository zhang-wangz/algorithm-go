package two_sequennce

// https://leetcode.cn/problems/longest-common-subsequence/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
// 一个字符串的 子序列 是指这样一个新的字符串：
// 它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

func longestCommonSubsequence(a string, b string) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(a)][len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
