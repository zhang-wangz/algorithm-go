package main

func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	gx := make([][]int, m+1)
	gy := make([][]int, m+1)
	for i := range gx {
		gx[i] = make([]int, n+1)
		gy[i] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			gx[i+1][j+1] = gx[i+1][j] + v
			gy[i+1][j+1] = gy[i][j+1] + v
		}
	}
	mx := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for j1 := n - 1; j1 >= j; j1-- {
				l := j1 - j + 1
				if l <= mx {
					break
				}
				if i+j1-j >= len(gx) || i+j1-j >= len(gy) {
					continue
				}
				if gx[i][j1]-gx[i][j-1] != l {
					continue
				}
				if gx[i+(j1-j)][j1]-gx[i+(j1-j)][j-1] != l {
					continue
				}
				if gy[i+(j1-j)][j]-gy[i-1][j] != l {
					continue
				}
				if gy[i+(j1-j)][j1]-gy[i-1][j1] != l {
					continue
				}
				mx = l
			}
		}
	}
	return mx * mx
}
