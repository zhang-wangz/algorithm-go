package main

// https://leetcode.cn/problems/omKAoA/

func minCut(s string) int {
	n := len(s)
	// 提前dp, 空间换时间
	mk := make([][]bool, n)
	for f := range mk {
		mk[f] = make([]bool, n)
		for j := range mk[f] {
			mk[f][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			mk[i][j] = s[i] == s[j] && mk[i+1][j-1]
		}
	}
	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i <= len(s); i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if mk[j][i-1] {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n]
}
