package solution

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 找出一块陆地
			// dfs标记当前块所属的整块陆地
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				cnt++
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, i, j int) int {
	if i > len(grid)-1 || j > len(grid[0])-1 || i < 0 || j < 0 {
		return 0
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1)
	}
	return 0
}
