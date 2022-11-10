package main

type pair struct{ x, y int }

func exist(board [][]byte, word string) bool {
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range dirs {
			if x, y := i+dir.x, j+dir.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] {
				if check(x, y, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
