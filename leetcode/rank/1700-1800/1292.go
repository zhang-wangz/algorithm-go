package main

// 1292
// 矩阵前缀和，i+1,j + i,j+1 - i,j + gij
func maxSideLength(mat [][]int, threshold int) int {
	maxsize := 0
	m, n := len(mat), len(mat[0])
	presum := make([][]int, m+1)
	for i, _ := range presum {
		presum[i] = make([]int, n+1)
	}
	for i, row := range mat {
		for j := range row {
			// 统计mat[0][0]到mat[i][j]的矩阵和
			// mat[0][0] 是 0
			presum[i+1][j+1] = presum[i][j+1] + presum[i+1][j] - presum[i][j] + mat[i][j]
			side := maxsize + 1
			for i+1-side >= 0 && j+1-side >= 0 && squareSum(presum, i, j, side) <= threshold {
				maxsize = side
				side++
			}
		}
	}
	return maxsize
}
func squareSum(presum [][]int, i, j, side int) int {
	return presum[i+1][j+1] - presum[i+1-side][j+1] - presum[i+1][j+1-side] + presum[i+1-side][j+1-side]
}
