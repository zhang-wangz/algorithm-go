package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	t := len(nums) - k
	quickSelect(nums, 0, len(nums)-1, t)
	return nums[t]
}

func quickSelect(nums []int, start, end int, idx int) {
	if len(nums) == 0 {
		return
	}
	p := rand.Intn(end-start+1) + start
	nums[p], nums[end] = nums[end], nums[p]
	p = nums[end]
	i, j := start, start
	for ; j < end; j++ {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	if i < idx {
		quickSelect(nums, i+1, end, idx)
	} else if i > idx {
		quickSelect(nums, start, i-1, idx)
	} else {
		return
	}
}
