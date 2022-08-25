package main

import "sort"

func findDiagonalOrder(nums [][]int) (ans []int) {
	type pair struct {
		ij, j, v int
	}
	arr := []pair{}
	for i, row := range nums {
		for j := range row {
			arr = append(arr, pair{i + j, j, row[j]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].ij != arr[j].ij {
			return arr[i].j < arr[j].j
		}
		return arr[i].ij < arr[j].ij
	})
	for _, p := range arr {
		ans = append(ans, p.v)
	}
	return
}

// 模拟超时
// 从左下到右上
func findDiagonalOrder2(nums [][]int) (ans []int) {
	m := len(nums)
	n := len(nums[0])
	for i := range nums {
		if len(nums[i]) > n {
			n = len(nums[i])
		}
	}
	for i := 0; i < m; i++ {
		l, r := i, 0
		for l >= 0 {
			if r < len(nums[l]) {
				ans = append(ans, nums[l][r])
			}
			l -= 1
			r += 1
		}
	}
	for i := 1; i < n; i++ {
		l, r := m-1, i
		for l >= 0 {
			if l < len(nums[l]) && r < len(nums[l]) {
				ans = append(ans, nums[l][r])
			}
			l -= 1
			r += 1
		}
	}
	return
}
