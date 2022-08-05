package main

import (
	"fmt"
)

// https://leetcode.cn/problems/NMTYve/
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var s string
	fmt.Scan(&s)
	mat := make([][]byte, 2*n+10)
	for i := range mat {
		mat[i] = make([]byte, 2*n+10)
		for j := range mat[i] {
			mat[i][j] = '.'
		}
	}
	up, down := -1<<63, 1<<63-1
	money := 0
	for i, c := range s {
		if c == '+' {
			up = max(money+n, up)
			down = min(money+n, down)
			mat[money+n][i] = '/'
			money++
		} else if c == '-' {
			money--
			up = max(money+n, up)
			down = min(money+n, down)
			mat[money+n][i] = '\\'
		} else {
			up = max(money+n, up)
			down = min(money+n, down)
			mat[money+n][i] = '-'
		}
	}

	for i := up; i >= down; i-- {
		for j := 0; j < n; j++ {
			fmt.Print(string(mat[i][j]))
		}
		fmt.Println()
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
