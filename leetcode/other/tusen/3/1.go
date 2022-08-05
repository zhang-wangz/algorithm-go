package main

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var subBox [3][3][9]int
	for i := range board {
		for j := range board[i] {
			c := board[i][j]
			if c != '.' {
				idx := c - '1'
				row[i][idx]++
				col[i][idx]++
				subBox[i/3][j/3][idx]++
				if row[i][idx] > 1 || col[j][idx] > 1 || subBox[i/3][j/3][idx] > 1 {
					return false
				}
			}
		}
	}
	return true
}
