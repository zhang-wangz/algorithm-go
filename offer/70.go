package main

// 非规则数的左边，偶数与+1的值相等，奇数与-1的值相等
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	// 右移的时候
	for l < r {
		mid := l + (r-l)>>1
		// mid为奇数的时候，mid^1为mid-1
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
