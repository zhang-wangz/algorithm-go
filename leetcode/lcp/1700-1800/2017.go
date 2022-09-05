package main

import "math"

// 2017 网格游戏
// 枚举拐弯的位置
func gridGame(grid [][]int) int64 {
	ans := math.MaxInt64
	l0 := 0
	for _, v := range grid[0] {
		l0 += v
	}
	l1 := 0
	for i, v := range grid[0] {
		l0 -= v
		ans = min(ans, max(l0, l1))
		l1 += grid[1][i]
	}
	return int64(ans)
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
