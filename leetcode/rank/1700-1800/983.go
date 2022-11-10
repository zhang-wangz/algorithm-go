package main

// 983
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	lastday, firstday := days[n-1], days[0]
	dp := make([]int, lastday+32)
	for d, i := lastday, n-1; d >= firstday; d-- {
		if days[i] == d {
			dp[d] = min(min(dp[d+1]+costs[0], dp[d+7]+costs[1]), dp[d+30]+costs[2])
			i--
		} else {
			dp[d] = dp[d+1]
		}
	}
	return dp[firstday]
}
