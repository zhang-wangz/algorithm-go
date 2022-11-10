package main

import "sort"

func countExcellentPairs(nums []int, k int) int64 {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		for j := i; j < len(nums); j++ {

		}
	}
	return 0
}

func compute(i, j int) int {
	cnt := 0
	a := i & j
	b := i | j
	for a != 0 {
		a &= a - 1
		cnt++
	}
	for b != 0 {
		b &= b - 1
		cnt++
	}
	return cnt
}
