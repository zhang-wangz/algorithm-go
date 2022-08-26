package main

import "math"

// 1262
func maxSumDivThree(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	return maxSumDivK(nums, 3)
}
func maxSumDivK(nums []int, k int) (ans int) {
	dp := make([]int, k)
	nextdp := make([]int, k)
	for i := range dp {
		dp[i] = math.MinInt
	}
	dp[0] = 0
	for _, num := range nums {
		for i := 0; i < k; i++ {
			nextdp[(num+i)%k] = max(dp[(num+i)%k], dp[i]+num)
		}
		for i := range dp {
			dp[i] = nextdp[i]
			nextdp[i] = 0
		}
	}
	return dp[0]
}

// dfs tle
func maxSumDivThree2(nums []int) (ans int) {
	sum := 0
	n := len(nums)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			return
		}
		if sum%3 == 0 {
			if sum > ans {
				ans = sum
			}
		}
		for i := idx; i < n; i++ {
			sum += nums[i]
			dfs(i + 1)
			sum -= nums[i]
		}
	}
	dfs(0)
	return
}
