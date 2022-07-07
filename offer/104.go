package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// dp[i]表示选取元素等于i的方案数
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			// 只要比i小都可以加进去
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}
