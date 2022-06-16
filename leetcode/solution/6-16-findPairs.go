package solution

import "sort"

// https://leetcode.cn/problems/k-diff-pairs-in-an-array/
// 给定一个整数数组和一个整数k，你需要在数组里找到 不同的k-diff 数对，并返回不同的 k-diff 数对 的数目。
// 这里将k-diff数对定义为一个整数对 (nums[i], nums[j])，并满足下述全部条件：
// 0 <= i < j < nums.length
// |nums[i] - nums[j]| == k
// 注意，|val| 表示 val 的绝对值。

// 二分，排序
func FindPairs(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := nums[i] + k
		tmp := -1
		l, r := i+1, len(nums)
		for l < r {
			mid := l + (r-l)>>1
			if mid > t {
				r = mid
			} else if mid < t {
				l = mid + 1
			} else {
				tmp = mid
				break
			}
		}
		if tmp != -1 && nums[tmp] == t {
			ans++
		}
	}
	return ans
}

// hashmap
func findPairs2(nums []int, k int) int {
	mp := map[int]bool{}
	ans := map[int]bool{}
	// 只存放x
	for _, e := range nums {
		if _, ok := mp[e-k]; ok {
			ans[e-k] = true
		}
		if _, ok := mp[e+k]; ok {
			ans[e] = true
		}
		mp[e] = true
	}
	return len(ans)
}

// 双指针+sort
func findPairs3(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0

	for i, j := 0, 0; i < len(nums)-1 && j < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < len(nums) && j <= i {
			j++
		}
		for j < len(nums) && nums[j]-nums[i] < k {
			j++
		}

		if j < len(nums) && nums[j]-nums[i] == k {
			ans++
		}
	}
	return ans
}
