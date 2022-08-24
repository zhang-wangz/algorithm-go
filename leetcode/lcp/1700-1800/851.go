package main

// 求>=当前i的数字中quiet值最小的人
// 求当前父节点中quiet最小的i
// richer是有向图
// 拓扑排序
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	in := make([]int, n)
	for _, e := range richer {
		g[e[0]] = append(g[e[0]], e[1])
		in[e[1]]++
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	q := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		nd := q[0]
		q = q[1:]
		// 遍历子节点
		for _, idx := range g[nd] {
			if quiet[ans[nd]] < quiet[ans[idx]] {
				ans[idx] = ans[nd]
			}
			if in[idx]--; in[idx] == 0 {
				q = append(q, idx)
			}
		}
	}
	return ans
}
