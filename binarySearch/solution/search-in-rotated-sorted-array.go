package solution

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 你可以假设数组中不存在重复的元素。
func search3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
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
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
