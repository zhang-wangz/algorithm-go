package _1_60

import "math"

// https://leetcode.cn/problems/7WqeDu/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		id := getId(nums[i], t+1)
		if _, ok := mp[id]; ok {
			return true
		}
		if y, ok := mp[id-1]; ok && int(math.Abs(float64(nums[i]-y))) <= t {
			return true
		}
		if y, ok := mp[id+1]; ok && int(math.Abs(float64(nums[i]-y))) <= t {
			return true
		}
		mp[id] = nums[i]
		if i >= k {
			delete(mp, getId(nums[i-k], t+1))
		}
	}
	return false
}

// 例如t=9， 0-9为一组，-1～-10也是一组，所以+1为0- -9，一组后-1即可
func getId(num, cap int) int {
	if num >= 0 {
		return num / cap
	}
	return (num+1)/cap - 1
}
