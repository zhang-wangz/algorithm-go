package solution

import "sort"

// https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/
// 给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。
// 同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子。
// 请你返回使用这块毯子，最多 可以盖住多少块瓷砖。
// 未解决
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	lstT := make([][]int, 0)
	for _, v := range tiles {
		lstT = append(lstT, []int{v[0], v[1], v[1] - v[0] + 1})
	}
	sort.Slice(lstT, func(i, j int) bool {
		return lstT[i][0] < lstT[j][0]
	})
	i := 0
	j := i + carpetLen - 1
	maxS := 0
	if len(lstT) == 1 && lstT[0][2] < carpetLen {
		return lstT[0][2]
	}
	for j < lstT[len(lstT)-1][1] {
		maxS = max(maxS, getLen(i, j, lstT))
		i++
		j++
	}
	return maxS
}

func getLen(i int, j int, t [][]int) int {
	l := 0
	if i >= t[0][0]-1 && j < t[0][1]-1 {
		l += j - i + 1
	} else if j > t[0][1]-1 {
		l += t[0][1] - 1 - i + 1
	}
	for k := 1; k < len(t); k++ {
		if i <= t[k][0]-1 && j >= t[k][0]-1 && j <= t[k][1]-1 {
			l += j - (t[k][0] - 1) + 1
		} else if i >= t[k][0]-1 && j < t[k][1]-1 {
			l += j - i + 1
		} else {
			l += t[0][1] - 1 - i + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
