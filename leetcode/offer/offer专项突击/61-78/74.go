package _1_78

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	if len(intervals) == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		l := intervals[i][0]
		r := intervals[i][1]
		if len(ans) == 0 || ans[len(ans)-1][1] < l {
			ans = append(ans, []int{l, r})
		} else {
			ans[len(ans)-1][1] = int(math.Max(float64(ans[len(ans)-1][1]), float64(r)))
		}
	}
	return ans
}
