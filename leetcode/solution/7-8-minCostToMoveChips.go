package solution

import "math"

// https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, n := range position {
		if n&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	return int(math.Min(float64(odd), float64(even)))
}

//func abs(a, b int) int {
//	if a > b {
//		return a - b
//	} else {
//		return b - a
//	}
//}
