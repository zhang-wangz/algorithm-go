package main

import "sort"

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	for i := 0; i < n; i++ {
		left, right := intervals[i][0], intervals[i][1]
		for ; i+1 < n && right >= intervals[i+1][0]; i++ {
			r := intervals[i+1][1]
			if r > right {
				right = r
			}
		}
		ans = append(ans, []int{left, right})
	}
	return
}
