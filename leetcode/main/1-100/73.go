package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := false, false
	// 第一行
	for _, c := range matrix[0] {
		if c == 0 {
			row = true
			break
		}
	}
	// 第一列
	for _, c := range matrix {
		if c[0] == 0 {
			col = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for j, _ := range matrix[0] {
			matrix[0][j] = 0
		}
	}

	if col {
		for _, row := range matrix {
			row[0] = 0
		}
	}
}
