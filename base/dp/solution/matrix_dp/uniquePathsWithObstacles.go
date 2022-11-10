package matrix_dp

// https://leetcode.cn/problems/unique-paths-ii/
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//  网格中的障碍物和空位置分别用 1 和 0 来表示。

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	res := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if res[i] == nil {
				res[i] = make([]int, len(obstacleGrid[i]))
			}
		}
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		res[0][i] = 1
	}

	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		res[i][0] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}

		}
	}
	return res[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
