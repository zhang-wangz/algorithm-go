package coinandbag

// https://leetcode.cn/problems/coin-change/
// dp 设置amount最少需要多少硬币
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			// 如果j比当前硬币数额大，减去之后的dp+1就是现在的数额
			if j-coins[i] >= 0 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
