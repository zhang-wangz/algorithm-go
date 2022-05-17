package solution

// 动态规划和 DFS 区别
// 二叉树 子问题是没有交集，所以大部分二叉树都用递归或者分治法，即 DFS，就可以解决
// 像 triangle 这种是有重复走的情况，子问题是有交集，所以可以用动态规划来解决

// 金字塔规划

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	res := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if res[i] == nil {
				res[i] = make([]int, len(triangle[i]))
			}
			res[i][j] = triangle[i][j]
		}
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[i][j] = min(res[i+1][j], res[i+1][j+1]) + triangle[i][j]
		}
	}
	return res[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
