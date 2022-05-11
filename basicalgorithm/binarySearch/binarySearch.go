package binarySearch

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	//  l = r+1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
			// [mid+1, r]
		} else if nums[mid] < target {
			l = mid + 1
			// [l, mid-1]
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
}

func binarySearch2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)
	// l == r
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
			// [mid+1, r)
		} else if nums[mid] < target {
			l = mid + 1
			// [l, mid)
		} else if nums[mid] > target {
			r = mid
		}
	}
	// 偏右搜索，终止时l == r，需要判断是否越界
	if l != len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func binarySearch3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
			//（mid, r)
		} else if nums[mid] < target {
			l = mid
			// (l, mid)
		} else {
			r = mid
		}
	}
	// 左右邻居
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
