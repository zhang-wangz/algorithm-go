package main

func minSideJumps(obstacles []int) int {
	dp := [3]int{}
	dp[0], dp[2] = 1, 1
	for i := 1; i < len(obstacles); i++ {
		ob := obstacles[i]
		pre0 := dp[0]
		pre1 := dp[1]
		pre2 := dp[2]
		for i := range dp {
			dp[i] = 5 * 10e5
		}
		if ob != 1 {
			dp[0] = pre0
		}
		if ob != 2 {
			dp[1] = pre1
		}
		if ob != 3 {
			dp[2] = pre2
		}

		if ob != 1 {
			dp[0] = min(dp[0], min(dp[1], dp[2])+1)
		}
		if ob != 2 {
			dp[1] = min(dp[1], min(dp[0], dp[2])+1)
		}
		if ob != 3 {
			dp[2] = min(dp[2], min(dp[0], dp[1])+1)
		}
	}
	ans := min(dp[0], min(dp[1], dp[2]))
	if ans == 5*10e5 {
		return -1
	}
	return ans
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
