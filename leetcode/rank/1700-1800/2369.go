package main

// dfs valid
func validPartition(nums []int) bool {
	n := len(nums)
	// 表示nums[0]到nums[i]的数字能否有效划分
	dp := make([]bool, n+1)
	// 空的肯定可以
	dp[0] = true
	for i, v := range nums {
		if i > 0 && dp[i-1] && v == nums[i-1] ||
			i > 1 && dp[i-2] && (v == nums[i-1] && v == nums[i-2]) ||
			v == nums[i-1]+1 && v == nums[i-2]+2 {
			dp[i+1] = true
		}
	}
	return dp[n]
}
