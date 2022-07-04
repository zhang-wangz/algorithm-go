package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	var dfs func(int, int)
	dfs = func(idx int, sum int) {
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
		}
		if sum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			// 删重复枝头
			if i-1 >= 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			used[i] = true
			dfs(i+1, sum)
			used[i] = false
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
