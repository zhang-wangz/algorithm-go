package solution

import "sort"

// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
// 给你一个整数数组 nums ，该数组具有以下属性：
//
// nums.length == 2 * n.
// nums 包含 n + 1 个 不同的 元素
// nums 中恰有一个元素重复 n 次
// 找出并返回重复了 n 次的那个元素。

func repeatedNTimes(nums []int) int {
	sort.Ints(nums)
	mid := (len(nums) - 1) / 2
	if mid-1 >= 0 && nums[mid-1] == nums[mid] {
		return nums[mid]
	} else {
		return nums[mid+1]
	}
}
