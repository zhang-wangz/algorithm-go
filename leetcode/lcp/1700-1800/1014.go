package main

// 1014 最佳观光组合
func maxScoreSightseeingPair(values []int) (ans int) {
	mx := values[0] + 0
	for i := 1; i < len(values); i++ {
		if mx+values[i]-i > ans {
			ans = mx + values[i] - i
		}
		// 维护当前mx
		mx = max(mx, values[i]+i)
	}
	return
}
