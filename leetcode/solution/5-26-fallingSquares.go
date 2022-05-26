package solution

// https://leetcode.cn/problems/falling-squares/
// 在无限长的数轴（即 x 轴）上，我们根据给定的顺序放置对应的正方形方块。
// 第 i 个掉落的方块（positions[i] = (left, side_length)）是正方形，
// 其中 left 表示该方块最左边的点位置(positions[i][0])，side_length 表示该方块的边长(positions[i][1])。
// 每个方块的底部边缘平行于数轴（即 x 轴），并且从一个比目前所有的落地方块更高的高度掉落而下。
// 在上一个方块结束掉落，并保持静止后，才开始掉落新方块。
// 方块的底边具有非常大的粘性，并将保持固定在它们所接触的任何长度表面上（无论是数轴还是其他方块）。
// 邻接掉落的边不会过早地粘合在一起，因为只有底边才具有粘性。
// 返回一个堆叠高度列表 ans 。每一个堆叠高度 ans[i]
// 表示在通过 positions[0], positions[1], ..., positions[i] 表示的方块掉落结束后，目前所有已经落稳的方块堆叠的最高高度。

// 暴力解法
func fallingSquares(positions [][]int) []int {
	res := make([]int, 0)
	old := make([][]int, 0)
	maxH := 0
	for i := 0; i < len(positions); i++ {
		start := positions[i][0]
		end := positions[i][0] + positions[i][1] - 1
		height := positions[i][1]
		hadd := 0
		for _, v := range old {
			if v[0] > end || v[1] < start {
				continue
			}
			hadd = max(hadd, v[2])
		}
		height = height + hadd
		maxH = max(height, maxH)
		old = append(old, []int{start, end, height})
		res = append(res, maxH)
	}
	return res
}

func getHash(hash map[int]int, i int) int {
	if v, ok := hash[i]; ok {
		return v
	} else {
		return -1
	}
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
