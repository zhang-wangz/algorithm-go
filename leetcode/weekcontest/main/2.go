package main

func equalPairs(grid [][]int) int {
	n := len(grid)
	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if tmp[i] == nil {
				tmp[i] = make([]int, n)
			}
			tmp[i][j] = grid[j][i]
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if grid[i][k] == tmp[j][k] {
					if k == n-1 {
						ans++
					}
				} else {
					break
				}
			}
		}
	}
	return ans
}
