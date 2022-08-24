package main

// 787
// bfs做吧
type pair struct{ idx, price, k int }

const INF = 0x3f3f3f3f

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	g := make([][]int, n)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = INF
	}
	for i := range g {
		g[i] = make([]int, n)
	}
	for _, e := range flights {
		g[e[0]][e[1]] = e[2]
	}
	q := []pair{{src, 0, 0}}
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		for i := range g[n.idx] {
			if g[n.idx][i] != 0 && n.price+g[n.idx][i] > dis[i] && n.k+1 <= k+1 {
				dis[i] = g[n.idx][i] + n.price
				q = append(q, pair{i, dis[i], n.k + 1})
			}
		}
	}
	if dis[dst] == INF {
		return -1
	}
	return dis[dst]
}
