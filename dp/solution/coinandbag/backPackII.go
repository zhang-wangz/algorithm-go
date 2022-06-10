package coinandbag

// https://www.lintcode.com/problem/125/
// 有 n 个物品和一个大小为 m 的背包.
// 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值. 问最多能装入背包的总价值是多大?
// 前i个物品在j大小的背包里能放的最大价值
func backPackII(m int, A []int, V []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			if j-A[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 在不放该物品前的最大价值和放物品后加上该物品的价值之间取最大值
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}
	return dp[len(A)][m]
}
