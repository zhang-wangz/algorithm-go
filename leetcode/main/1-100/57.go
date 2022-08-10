package main

import "sort"

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	intervals = append(intervals, newInterval)
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

func insert2(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if right < interval[0] {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if left > interval[1] {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
