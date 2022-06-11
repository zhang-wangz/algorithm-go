package solution

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
// 如果一个二进制字符串，是以一些 0（可能没有 0）后面跟着一些 1（也可能没有 1）的形式组成的，那么该字符串是 单调递增 的。
// 给你一个二进制字符串 s，你可以将任何 0 翻转为 1 或者将 1 翻转为 0 。
// 返回使 s 单调递增的最小翻转次数。

func MinFlipsMonoIncr(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= len(s); i++ {
		dp[i][0] = dp[i-1][0] + int(s[i-1]-'0')
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + (int(s[i-1]-'0') ^ 1)
	}
	return min(dp[len(s)][0], dp[len(s)][1])
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
