package sort

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
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}
	// 把end位置换到i，并且返回中间标志值
	tmp := nums[i]
	nums[i] = nums[end]
	nums[end] = tmp
	return i
}
