package main

import (
	"math"
	"sort"
)

// 枚举对角线
func minAreaRect(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	type pair struct{ i, j int }
	m := map[pair]bool{}
	ans = math.MaxInt
	for _, p := range points {
		i, j := p[0], p[1]
		for k := range m {
			x, y := k.i, k.j
			if m[pair{i, y}] && m[pair{x, j}] {
				ans = min(ans, (j-y)*(i-x))
			}
		}
		m[pair{i, j}] = true
	}
	if ans == math.MaxInt {
		ans = 0
	}
	return
}
