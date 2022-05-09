package solution

// 多源起点bfs  转换为0到1的距离
// https://leetcode.cn/problems/01-matrix/
func UpdateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	m := len(mat)
	n := len(mat[0])
	visit := make([][]int, 0)
	q := make([][]int, 0)
	for i := 0; i < m; i++ {
		visit = append(visit, make([]int, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
				visit[i][j] = 1
			} else {
				mat[i][j] = -1
			}
		}
	}
	step := 1
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			p := q[0]
			q = q[1:]
			direct := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
			for v := 0; v < len(direct); v++ {
				x := p[0] + direct[v][0]
				y := p[1] + direct[v][1]

				if x > 0 && y > 0 && x < m && y < n && mat[x][y] == -1 && visit[x][y] == 0 {
					mat[x][y] = step
					visit[x][y] = 1
					q = append(q, []int{x, y})
				}
			}
		}
		step += 1
	}
	return mat
}
