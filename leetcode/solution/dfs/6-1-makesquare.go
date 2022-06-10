package dfs

import "sort"

// 火柴堆正方形
func Makesquare(matchsticks []int) bool {
	sum := 0
	for _, num := range matchsticks {
		sum += num
	}
	if sum%4 != 0 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	edge := make([]int, 4)
	var dfs func(matchsticks []int, start int, edge []int) bool
	dfs = func(matchsticks []int, start int, edge []int) bool {
		if start == len(matchsticks) {
			return true
		}
		for i := 0; i < len(edge); i++ {
			edge[i] += matchsticks[start]
			if edge[i] <= sum/4 && dfs(matchsticks, start+1, edge) {
				return true
			}
			edge[i] -= matchsticks[start]
		}
		return false
	}
	return dfs(matchsticks, 0, edge)
}
