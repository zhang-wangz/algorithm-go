package solution

// 与1的区别是元素可能存在重复元素
func findMin2(nums []int) int {
	// 把最后一个作为target
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l+1 < r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}
		mid := l + (r-l)/2
		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid
		}
	}
	if nums[l] > nums[r] {
		return nums[r]
	}
	return nums[l]
}
