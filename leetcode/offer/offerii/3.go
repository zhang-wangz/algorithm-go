package main
// https://leetcode.cn/problems/find-the-duplicate-number/description/
func findRepeatNumber(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] == 2 {
			return nums[i]
		}
	}
	return -1
}
