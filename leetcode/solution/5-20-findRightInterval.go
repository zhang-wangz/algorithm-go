package solution

import "sort"

// https://leetcode.cn/problems/find-right-interval/
// 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
// 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
// 返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1

// 查找换成2分查找就不超时了， 感觉可以用回溯做，待解决
func FindRightInterval(intervals [][]int) []int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return []int{-1}
	}
	m := make(map[int]int)
	res := make([]int, len(intervals))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for i := 0; i < len(intervals); i++ {
		m[intervals[i][0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		l := 0
		r := len(intervals) - 1
		for l+1 < r {
			mid := l + (r-l)/2
			if intervals[mid][0] >= intervals[i][1] {
				r = mid
			} else {
				l = mid
			}
		}
		if intervals[l][0] >= intervals[i][1] {
			res[m[intervals[i][0]]] = m[intervals[l][0]]
		} else if intervals[r][0] >= intervals[i][1] {
			res[m[intervals[i][0]]] = m[intervals[r][0]]
		}
	}
	return res
}

// 回溯
//func findRightInterval(intervals [][]int) []int {
//	if len(intervals) == 0 || len(intervals) == 1 {
//		return []int{-1}
//	}
//	res := make([]int, len(intervals))
//	for i := 0; i < len(res); i++ {
//		res[i] = -1
//	}
//	back(res, intervals)
//	return res
//}
//
//func back(res []int, intervals [][]int){
//
//}
