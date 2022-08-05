package main

import "math"

func maxProfit(prices []int) int {
	ans := 0
	ml := math.MaxInt
	for _, p := range prices {
		if p < ml {
			ml = p
		}
		ans = int(math.Max(float64(p-ml), float64(ans)))
	}
	return ans
}
