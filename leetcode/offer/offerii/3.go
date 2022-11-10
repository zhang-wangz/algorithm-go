package main

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
