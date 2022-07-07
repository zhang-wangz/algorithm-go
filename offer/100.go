package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	res := make([][]int, len(triangle))
	for f := range res {
		res[f] = make([]int, len(triangle[f]))
		for j := range res[f] {
			res[f][j] = triangle[f][j]
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[i][j] = min(res[i+1][j], res[i+1][j+1]) + triangle[i][j]
		}
	}
	return res[0][0]
}
