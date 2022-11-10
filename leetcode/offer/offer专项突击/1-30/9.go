package offer专项突击_1_30

// 给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

// 每次增加的子数组的数量就是 当前子数组的长度
func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	right := 0
	mul := 1
	ans := 0
	for right < len(nums) {
		t := nums[right]
		right++
		mul *= t
		for mul >= k && left < right {
			t := nums[left]
			left++
			mul /= t
		}
		ans += right - left
	}
	return ans
}
