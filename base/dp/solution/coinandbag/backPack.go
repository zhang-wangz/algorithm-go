package coinandbag

// https://www.lintcode.com/problem/92/
// 在 n 个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为 m，每个物品的大小为 A[i]
// dp[i][j] 表示前i个的物品装进j大小的背包最多能装多满
func backPack(m int, A []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			// 如果j大小的背包放不下第i个物品
			// 就装的和i-1个物品能装的一样
			if j-A[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+A[i-1])
			}
		}
	}
	return dp[len(A)][m]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
