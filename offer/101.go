package main

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}

	//dp := make([][]bool, n)
	//for i := 0; i < len(dp); i++ {
	//	dp[i] = make([]bool, target+1)
	//	dp[i][0] = true
	//}
	dp := make([]bool, target+1)
	dp[0] = true
	// 因i-1只和i有关所以可以优化
	for i := 1; i < n; i++ {
		//for j := 1; j <= target; j++ {
		// 需要从大到小，不然dp[j]无法构成一致
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]

			//if j >= nums[i] {
			//	dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			//} else {
			//	dp[i][j] = dp[i-1][j]
			//}
		}
	}
	return dp[target]
	//return dp[n-1][target]
}
