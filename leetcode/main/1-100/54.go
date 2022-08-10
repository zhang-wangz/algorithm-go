package main

// https://leetcode.cn/problems/spiral-matrix/
// 右下左上
var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return
	}
	tot := m * n
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	x, y := 0, 0
	idx := 0
	for k := 0; k < tot; k++ {
		vis[x][y] = true
		ans = append(ans, matrix[x][y])
		i, j := x+dirs[idx][0], y+dirs[idx][1]
		if i < 0 || i >= m || j < 0 || j >= n || vis[i][j] {
			idx = (idx + 1) % 4
		}
		i, j = x+dirs[idx][0], y+dirs[idx][1]
		x, y = i, j
	}
	return
}
