package main

func minArray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}
		mid := l + (r-l)>>1
		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
