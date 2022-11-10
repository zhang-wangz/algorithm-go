package main

// 1024
func videoStitching(clips [][]int, time int) int {
	const inf = 0x3f3f3f3f
	dp := make([]int, time+1)
	for i := range dp {
		dp[i] = inf
	}
	// [0,i）的dp所需要的最小合并次数
	dp[0] = 0
	for i := 1; i <= time; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[time] == inf {
		return -1
	}
	return dp[time]
}
