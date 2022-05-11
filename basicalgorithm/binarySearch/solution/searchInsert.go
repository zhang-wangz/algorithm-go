package solution

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	r := len(nums) - 1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid
		}
	}
	if nums[l] >= target {
		return l
	} else if nums[r] >= target {
		return r
	} else {
		return r + 1
	}
}
