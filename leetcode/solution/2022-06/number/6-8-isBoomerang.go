package number

// https://leetcode.cn/problems/valid-boomerang/
// 给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，
// 如果这些点构成一个 回旋镖 则返回 true 。
// 回旋镖 定义为一组三个点，这些点 各不相同 且 不在一条直线上 。

func isBoomerang(points [][]int) bool {
	k1 := []int{points[0][0] - points[1][0], points[0][1] - points[1][1]}
	k2 := []int{points[1][0] - points[2][0], points[1][1] - points[2][1]}
	return k2[0]*k1[1] != k1[0]*k2[1]
}
