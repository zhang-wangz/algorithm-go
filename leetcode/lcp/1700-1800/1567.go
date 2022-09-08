package main

// 1567 乘积为正数的最长子数组长度
func getMaxLen(nums []int) (ans int) {
	n := len(nums)
	pos, neg := make([]int, n), make([]int, n)
	if nums[0] < 0 {
		neg[0] = 1
	} else if nums[0] > 0 {
		pos[0] = 1
	}
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			pos[i] = pos[i-1] + 1
			if neg[i-1] > 0 {
				neg[i] = neg[i-1] + 1
			}
		} else if nums[i] < 0 {
			if neg[i-1] > 0 {
				pos[i] = neg[i-1] + 1
			}
			neg[i] = pos[i-1] + 1
		}
		ans = max(ans, pos[i])
	}
	return
}
