package dp_88_104

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func minCostClimbingStairs2(cost []int) int {
	pre, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, min(pre+cost[i-2], cur+cost[i-1])
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
