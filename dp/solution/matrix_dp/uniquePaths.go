package matrix_dp

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i] == nil {
				res[i] = make([]int, n)
			}
			res[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}
