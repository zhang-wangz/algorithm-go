package offer专项突击_1_30

import "math"

// https://leetcode.cn/problems/A1NYOS/
func findMaxLength(nums []int) int {
	// 存储第一次前缀和下标
	mp := map[int]int{0: -1}
	sum := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 是1就+1，是0就-1
		sum += num & 1
		sum -= num ^ 1
		if v, ok := mp[sum]; ok {
			ans = int(math.Max(float64(ans), float64(i-v)))
		} else {
			mp[sum] = i
		}
	}
	return ans
}
