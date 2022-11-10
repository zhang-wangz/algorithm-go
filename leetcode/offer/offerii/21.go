package main

func exchange(nums []int) []int {
	low, fast := 0, 0
	n := len(nums)
	for fast < n {
		if nums[fast]&1 == 1 {
			nums[fast], nums[low] = nums[low], nums[fast]
			low++
		}
		fast++
	}
	return nums
}
