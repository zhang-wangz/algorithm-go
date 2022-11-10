package sort

import "math/rand"

func QuickSelect(nums []int, start, end int, tidx int) {
	p := rand.Intn(end-start+1) + start
	nums[p], nums[end] = nums[end], nums[p]
	p = nums[end]
	i, j := start, start
	for j < end {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	if i > tidx {
		QuickSelect(nums, start, i-1, tidx)
	} else if i < tidx {
		QuickSelect(nums, i+1, end, tidx)
	} else {
		return
	}
}
