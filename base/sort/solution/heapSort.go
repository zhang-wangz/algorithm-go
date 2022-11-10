package solution

// 升序排列
func HeapSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	// 排一个最大堆, 找出最大的节点
	sink(nums, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sink(nums, i)
	}
	return nums
}

// 第一个非叶子节点是n/2-1
func sink(nums []int, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		left := 2*i + 1
		if left < len && nums[left] > nums[i] {
			nums[i], nums[left] = nums[left], nums[i]
		}
		right := 2*i + 2
		if right < len && nums[right] > nums[i] {
			nums[i], nums[right] = nums[right], nums[i]
		}
	}
}
