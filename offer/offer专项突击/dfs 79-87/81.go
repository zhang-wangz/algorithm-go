package dfs_79_87

import "sort"

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(sum int, idx int) {
		if idx == len(candidates) {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
		}
		if sum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(sum, i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}
