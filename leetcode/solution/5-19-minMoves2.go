package solution

import "sort"

// https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
// 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
// 在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。
func minMoves2(nums []int) int {
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	mid := l + (r-l)/2
	target := nums[mid]
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += abs(nums[i], target)
	}
	return cnt
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
