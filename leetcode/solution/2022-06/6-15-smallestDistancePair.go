package _022_06

import "sort"

// https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
// 数对 (a,b) 由整数 a 和 b 组成，其数对距离定义为 a 和 b 的绝对差值。
//
//给你一个整数数组 nums 和一个整数 k ，数对由 nums[i] 和 nums[j] 组成且满足 0 <= i < j < nums.length 。
//返回 所有数对距离中 第 k 小的数对距离。

func SmallestDistancePair2(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, 1000000+1
	for l < r {
		mid := l + (r-l)/2
		cnt := 0
		for i, j := 0, 0; i < len(nums); i++ {
			for j < len(nums) && nums[j]-nums[i] <= mid {
				j++
			}
			cnt += j - i - 1
		}
		// 获取符合要求的第一个值
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// TLE
// 简单写个堆排吧
func SmallestDistancePair(nums []int, k int) int {
	des := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			des = append(des, abs(nums[i], nums[j]))
		}
	}
	sink(des, len(des))
	//cnt:=0
	for i := len(des) - 1; i >= len(des)-k; i-- {
		des[0], des[i] = des[i], des[0]
		sink(des, i)
	}
	return des[len(des)-k]
}

func sink(nums []int, l int) {
	// 最后一个非叶子结点开始下沉
	i := l/2 - 1
	for i >= 0 {
		idxl := 2*i + 1
		idxr := 2*i + 2
		if idxl < l && nums[idxl] < nums[i] {
			nums[idxl], nums[i] = nums[i], nums[idxl]
		}

		if idxr < l && nums[idxr] < nums[i] {
			nums[idxr], nums[i] = nums[i], nums[idxr]
		}
		i--
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
