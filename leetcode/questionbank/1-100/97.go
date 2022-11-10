package main

// https://leetcode.cn/problems/interleaving-string/
// s3是否由s1和s2交错组成的
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, f := len(s1), len(s2), len(s3)
	if m+n != f {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || ((s3[p] == s1[i-1]) && dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || ((s3[p] == s2[j-1]) && dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
