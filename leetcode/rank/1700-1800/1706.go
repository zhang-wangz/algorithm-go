package main

func findBall(grid [][]int) (ans []int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		x, y := 0, i
		for y >= 0 && y < n && x >= 0 && x < m {
			if y < n-1 && grid[x][y] == 1 && 1 == grid[x][y+1] {
				x += 1
				y += 1
			} else if y > 0 && grid[x][y] == -1 && -1 == grid[x][y-1] {
				x += 1
				y -= 1
			} else {
				break
			}
		}
		if x == m {
			ans = append(ans, y)
		} else {
			ans = append(ans, -1)
		}
	}
	return
}
