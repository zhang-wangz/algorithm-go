package main

const (
	M = 6010
	N = 110
)
const INF = 0x3f3f3f3f

type edge struct {
	a, b, w int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func networkDelayTime(times [][]int, n int, k int) int {
	dist := [N]int{}
	es := [M]edge{}
	for t, time := range times {
		es[t] = edge{time[0], time[1], time[2]}
	}

	// bf
	for i := 0; i <= n; i++ {
		dist[i] = INF
	}
	dist[k] = 0
	for i := 1; i < n; i++ {
		flag := false
		for _, e := range es {
			a, b, w := e.a, e.b, e.w
			if dist[b] > dist[a]+w {
				dist[b] = dist[a] + w
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dist[i])
	}
	if ans == INF {
		return -1
	} else {
		return ans
	}
}
