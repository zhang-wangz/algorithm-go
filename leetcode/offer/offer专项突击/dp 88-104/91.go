package dp_88_104

//  https://leetcode.cn/problems/JEj789/

func minCost(costs [][]int) int {
	a := costs[0][0]
	b := costs[0][1]
	c := costs[0][2]
	for i := 1; i < len(costs); i++ {
		d := _min(b, c) + costs[i][0]
		e := _min(a, c) + costs[i][1]
		f := _min(a, b) + costs[i][2]
		a, b, c = d, e, f
	}
	return _min(a, _min(b, c))
}

func _min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
