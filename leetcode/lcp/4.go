package main

// https://leetcode.cn/problems/broken-board-dominoes/

func domino(n int, m int, broken [][]int) int {
	broke := map[int]map[int]int{}
	for _, b := range broken {
		if broke[b[0]] == nil {
			broke[b[0]] = map[int]int{}
		}
		broke[b[0]][b[1]] = 1
	}
	// todo
	return 0
}
