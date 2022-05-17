package matrix_dp

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j-1 < 0 && i-1 >= 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if j-1 >= 0 && i-1 < 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j-1 >= 0 && i-1 >= 0 {
				grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
