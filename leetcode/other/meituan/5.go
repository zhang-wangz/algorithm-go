package main

import "math"

func candy(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, c := range ratings {
		if i > 0 && c > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += int(math.Max(float64(left[i]), float64(right)))
	}
	return
}
