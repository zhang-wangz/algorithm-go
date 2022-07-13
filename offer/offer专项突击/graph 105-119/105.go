package graph_105_119

import "math"

type pair struct {
	i, j int
}

// stack 和queue互换一下就是dfs和bfs的交换
func maxAreaOfIsland(grid [][]int) (ans int) {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	q := make([]*pair, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := 0
			q = append(q, &pair{i, j})
			for len(q) != 0 {
				node := q[0]
				q = q[1:]
				ni, nj := node.i, node.j
				if ni < 0 || ni >= m || nj < 0 || nj >= n || grid[ni][nj] != 1 {
					continue
				}
				grid[ni][nj] = 0
				cur++
				dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
				for _, dir := range dirs {
					nexti, nextj := ni+dir[0], nj+dir[1]
					q = append(q, &pair{nexti, nextj})
				}
			}
			ans = int(math.Max(float64(ans), float64(cur)))
		}
	}
	return
}
