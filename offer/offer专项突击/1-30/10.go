package offer专项突击_1_30

// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
// 因为数据范围包括负数，所以不能用滑动窗口

func subarraySum(nums []int, k int) int {
	cnt := 0
	presum := 0
	mp := map[int]int{
		0: 1,
	}
	for i := 0; i < len(nums); i++ {
		presum += nums[i]
		cnt += mp[presum-k]
		mp[presum] = mp[presum] + 1
	}
	return cnt
}
