package main

func generateMatrix(n int) (ans [][]int) {
	var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	tot := n * n
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	x, y := 0, 0
	idx := 0
	cnt := 1
	for k := 0; k < tot; k++ {
		vis[x][y] = true
		ans[x][y] = cnt
		cnt++
		i, j := x+dirs[idx][0], y+dirs[idx][1]
		if i < 0 || i >= n || j < 0 || j >= n || vis[i][j] {
			idx = (idx + 1) % 4
		}
		i, j = x+dirs[idx][0], y+dirs[idx][1]
		x, y = i, j
	}
	return
}
