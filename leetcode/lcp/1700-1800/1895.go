package main

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowsum, colsum := make([][]int, m+1), make([][]int, m+1)
	line45, line135 := make([][]int, m+1), make([][]int, m+1)
	for i := range colsum {
		rowsum[i] = make([]int, n+1)
		colsum[i] = make([]int, n+1)
		line45[i] = make([]int, n+1)
		line135[i] = make([]int, n+1)
	}
	// 求各自的前缀和
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowsum[i][j+1] = rowsum[i][j] + grid[i][j]
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			colsum[i+1][j] = colsum[i][j] + grid[i][j]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			line45[i][j] = line45[i-1][j+1] + grid[i-1][j]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			line135[i][j] = line135[i-1][j-1] + grid[i-1][j-1]
		}
	}
	isValid := func(i, j, k int) bool {
		target := line45[i+k][j] - line45[i][j+k]
		for ii := i; ii < i+k; ii++ {
			sum := rowsum[ii][j+k] - rowsum[ii][j]
			if sum != target {
				return false
			}
		}
		for jj := j; jj < j+k; jj++ {
			sum := colsum[i+k][jj] - colsum[i][jj]
			if sum != target {
				return false
			}
		}
		return line135[i+k][j+k]-line135[i][j] == target
	}
	for k := min(m, n); k >= 2; k-- {
		for i := 0; i+k-1 < m; i++ {
			for j := 0; j+k-1 < n; j++ {
				if isValid(i, j, k) {
					return k
				}
			}
		}
	}
	return 1
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
