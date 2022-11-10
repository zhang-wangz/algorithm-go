package sequence

// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个位置.
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	r := make([]bool, len(nums))
	r[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if r[j] == true && nums[j]+j >= i {
				r[i] = true
			}
		}
	}
	return r[len(nums)-1]
}
