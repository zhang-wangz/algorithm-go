package _022_08

import "sort"

// https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order/
func minSubsequence(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	tot := 0
	for i, num := range nums {
		tot += num
		if tot > sum-tot {
			return nums[:i+1]
		}
	}
	return []int{}
}
