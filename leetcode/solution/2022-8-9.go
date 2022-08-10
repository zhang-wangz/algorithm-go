package main

import "math"

// https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/

func minStartValue(nums []int) int {
	tot := 0
	ml := 0
	for _, num := range nums {
		tot += num
		ml = int(math.Min(float64(ml), float64(tot)))
	}
	return 1 - ml
}
