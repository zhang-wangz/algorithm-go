package sort

import "math/rand"

// 分治思想快排
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if len(nums) <= 1 {
		return
	}
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
}

func partition(nums []int, start, end int) int {
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
	// 把end位置换到i，并且返回中间标志值
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
