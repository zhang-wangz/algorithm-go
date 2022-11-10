package dp_88_104

// 找出0-i中是否有存在总和为neg的存在
func findTargetSumWays2(nums []int, target int) (ans int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}

func findTargetSumWays(nums []int, target int) (ans int) {
	n := len(nums)
	path := 0
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n && path == target {
			ans++
			return
		}
		if idx == n {
			return
		}
		path += nums[idx]
		dfs(idx + 1)
		path -= nums[idx]
		path -= nums[idx]
		dfs(idx + 1)
		path += nums[idx]
	}
	dfs(0)
	return
}
