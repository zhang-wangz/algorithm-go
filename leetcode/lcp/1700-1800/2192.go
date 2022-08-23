package main

// 三叶姐的go版本～
func getAncestors(n int, edges [][]int) (ans [][]int) {
	in := make([]int, n)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		in[e[1]]++
	}
	pa := make([][]bool, n)
	for i := range pa {
		pa[i] = make([]bool, n)
	}
	q := make([]int, 0)
	for k, v := range in {
		if v == 0 {
			q = append(q, k)
		}
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		// 遍历n的子节点
		for _, v := range g[n] {
			for i := range pa[n] {
				// 如果n是i的子节点，v也是i的子节点
				if pa[n][i] {
					pa[v][i] = true
				}
			}
			// v是n的子节点
			pa[v][n] = true
			if in[v]--; in[v] == 0 {
				q = append(q, v)
			}
		}
	}
	ans = make([][]int, n)
	for i, row := range pa {
		for j := range row {
			if row[j] {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}

//func main() {
//	s := [][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}}
//	getAncestors(8,s)
//}
