package solution

func QuickSort(nums []int, start, end int) []int {
	quick := quickSort(nums, 0, len(nums)-1)
	return quick
}

func quickSort(nums []int, start, end int) []int {
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
	return nums
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	var i int  // 标记比p小的值的idx
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
