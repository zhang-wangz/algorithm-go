package offer

// https://leetcode.cn/problems/tvdfij/
func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	cnt := -1
	for i := 0; i < len(sum); i++ {
		if i != len(nums) && sum[i] == sum[len(nums)]-sum[i]-nums[i] {
			cnt = i
			break
		}
	}
	return cnt
}
