package dp_88_104

func rob(nums []int) int {
	pre := nums[0]
	if len(nums) == 1 {
		return pre
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	cur := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		next := max(nums[i]+pre, cur)
		pre = cur
		cur = next
	}
	return cur
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
