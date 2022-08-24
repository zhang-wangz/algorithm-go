package main

// 1911
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	dp := make([][2]int, n+1)
	dp[0][0] = 0
	dp[0][1] = -1
	for i := 1; i <= len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-nums[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i-1])
	}
	return int64(max(dp[n][0], dp[n][1]))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
