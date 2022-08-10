package main

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	path := []int{}
	sum := 0
	var dfs func(start int)
	vis := make([]bool, len(candidates))
	dfs = func(start int) {
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] && !vis[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			vis[i] = true
			dfs(i + 1)
			vis[i] = false
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
