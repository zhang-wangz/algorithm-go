package lcp

// https://leetcode.cn/problems/broken-board-dominoes/

// 二分模版
func domino(n int, m int, broken [][]int) int {
	edge := make([][]bool, 65)
	for i := 0; i < len(edge); i++ {
		edge[i] = make([]bool, 65)
	}
	vis := make([]bool, 65)
	match := make([]int, 65)
	for i := 0; i < len(match); i++ {
		match[i] = -1
	}
	// 总共n*m个点
	N := n * m
	var dfs func(u int) bool
	dfs = func(u int) bool {
		for i := 0; i < N; i++ {
			if edge[u][i] && !vis[i] {
				vis[i] = true
				if match[i] == -1 || dfs(match[i]) {
					match[i] = u
					return true
				}
			}
		}
		return false
	}

	// build graph
	// 0代表可以放
	graph := make([][]int, n)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
	}
	// 1代表是坏的不能放
	for i := 0; i < len(broken); i++ {
		graph[broken[i][0]][broken[i][1]] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 1 {
				continue
			}
			for _, dir := range dirs {
				dx, dy := i+dir[0], j+dir[1]
				if dx < 0 || dx >= n || dy < 0 || dy >= m || graph[dx][dy] == 1 {
					continue
				}
				if (i+j)%2 == 0 {
					// 只记录偶数向奇数的边
					edge[i*m+j][dx*m+dy] = true
				}
			}
		}
	}

	ans := 0
	// 因为计算是边的，所以就直接算N也不会有影响，只是浪费了空间
	for i := 0; i < N; i++ {
		vis = make([]bool, 65)
		// dfs每个男孩纸的配对
		// 若能配对，则匹配+1
		if dfs(i) {
			ans++
		}
	}
	return ans
}
