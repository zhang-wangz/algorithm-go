package _022_07

import "sort"

// https://leetcode.cn/problems/set-intersection-size-at-least-two/
// 先把区间进行排序，然后从后往前逐一进行判断，如果符合交集个数则跳过，否则从该区间取数字加入交集
func intersectionSizeTwo(intervals [][]int) int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	ans := 0
	// 交集定义
	vals := make([][]int, n)
	m := 2
	for i := n - 1; i >= 0; i-- {
		for j, k := intervals[i][0], len(vals[i]); k < m; k++ {
			ans++
			for p := i - 1; p >= 0 && intervals[p][1] >= j; p-- {
				vals[p] = append(vals[p], j)
			}
			j++
		}
	}
	return ans
}
