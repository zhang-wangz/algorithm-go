package main

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, n)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		memo[i][j]++
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
				memo[i][j] = max(memo[i][j], dfs(ni, nj)+1)
			}
		}
		return memo[i][j]
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
