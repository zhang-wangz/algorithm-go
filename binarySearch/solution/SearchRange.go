package solution

// https://www.lintcode.com/problem/61/
// 给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。 如果目标值不在数组中，则返回[-1, -1]
func SearchRange(a []int, target int) []int {
	if len(a) == 0 {
		return []int{-1, -1}
	}
	res := []int{-1, -1}
	l := 0
	r := len(a) - 1
	for l+1 < r {
		mid := l + (r-l)/2
		if a[mid] < target {
			l = mid
		} else if a[mid] > target {
			r = mid
		} else {
			r = mid
		}
	}
	if a[l] == target {
		res[0] = l
	} else if a[r] == target {
		res[0] = r
	} else {
		return []int{-1, -1}
	}

	l = 0
	r = len(a) - 1
	for l+1 < r {
		mid := l + (r-l)/2
		if a[mid] < target {
			l = mid
		} else if a[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	if a[r] == target {
		res[1] = r
	} else if a[l] == target {
		res[1] = l
	} else {
		return []int{-1, -1}
	}
	return res
}
