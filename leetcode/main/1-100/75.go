package main

// https://leetcode.cn/problems/sort-colors/
// 3指针
func sortColors(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	i := 0
	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
}
