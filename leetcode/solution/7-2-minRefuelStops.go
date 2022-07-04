package solution

import "math"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	// 加了i次油能够拥有的最大汽油数
	dp := make([]int, len(stations)+1)
	dp[0] = startFuel
	for i := 0; i < len(stations); i++ {
		for j := i; j >= 0; j-- {
			if dp[j] >= stations[i][0] {
				dp[j+1] = int(math.Max(float64(dp[j+1]), float64(dp[j]+stations[i][1])))
			}
		}
	}
	for i := 0; i <= len(stations); i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}
