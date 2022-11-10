package solution

// 1的元素重复版
func search4(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for l+1 < r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}
		if l+1 >= r {
			break
		}
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		// 左边有序
		if nums[l] < nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid
			} else {
				l = mid
			}
			// 右边有序
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return true
	}
	if nums[r] == target {
		return true
	}
	return false
}
