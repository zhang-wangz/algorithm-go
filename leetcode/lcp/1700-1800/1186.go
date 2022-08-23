package main

// 状态dp
// 由求最大子数组和转化而来，多了一个删除的状态
func maximumSum(arr []int) (ans int) {
	dp := make([][2]int, len(arr))
	dp[0][0] = arr[0]
	dp[0][1] = 0
	for i := 1; i < len(arr); i++ {
		dp[i][0] = max(0, dp[i-1][0]) + arr[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		ans = max(ans, max(dp[i][0], dp[i][1]))
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
