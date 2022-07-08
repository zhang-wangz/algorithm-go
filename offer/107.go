package main

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	ans := make([][]int, m)
	vis := make([][]bool, m)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, n)
		vis[i] = make([]bool, n)
	}
	q := make([]*pair, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, &pair{i, j})
				vis[i][j] = true
			} else {
				ans[i][j] = -1
			}
		}
	}
	// 多源起点的时候，从起点出发，不要去求每个结束点到源点的距离
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		i, j := node.i, node.j
		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !vis[ni][nj] {
				ans[ni][nj] = ans[i][j] + 1
				q = append(q, &pair{ni, nj})
				vis[ni][nj] = true
			}
		}
	}
	return ans
}
