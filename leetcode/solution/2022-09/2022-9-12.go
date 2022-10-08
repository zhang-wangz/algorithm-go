package main

import "sort"

// 1608  官方题解理解
func specialArray(nums []int) int {
	// 倒序排列，从大到小
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// i-1>=i 说明目前有i个数字>=i（0到i-1）
	// 后面i==n时说明有n个数字大于等于n，退出循环
	// 另一个则是第i个数后面的数<i，也满足判断
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}
