package dp

// https://leetcode.cn/problems/count-different-palindromic-subsequences/
// 给定一个字符串 s，返回 s 中不同的非空「回文子序列」个数 。
// 通过从 s 中删除 0 个或多个字符来获得子序列。
// 如果一个字符序列与它反转后的字符序列一致，那么它是「回文字符序列」。
// 如果有某个 i , 满足 ai != bi ，则两个序列 a1, a2, ... 和 b1, b2, ... 不同。
// 注意：结果可能很大，你需要对 109 + 7 取模 。
// hard dp

func CountPalindromicSubsequences(s string) int {
	mod := 1000000007
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				left := i + 1
				right := j - 1
				for left <= right && s[left] != s[i] {
					left++
				}
				for left <= right && s[right] != s[j] {
					right--
				}
				if left > right {
					// 没有重复
					dp[i][j] = 2*dp[i+1][j-1] + 2
				} else if left == right {
					// 1个重复
					dp[i][j] = 2*dp[i+1][j-1] + 1
				} else {
					// 2个及以上
					dp[i][j] = 2*dp[i+1][j-1] - dp[left+1][right-1]
				}
			} else {
				dp[i][j] = dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1]
			}
			// mod转移导致负数
			// dp[left+1][right-1]可能大于dp[i+1][j-1]*2
			if dp[i][j] >= 0 {
				dp[i][j] = dp[i][j] % mod
			} else {
				dp[i][j] = dp[i][j] + mod
			}
		}
	}
	return dp[0][n-1]
}
