package main

const (
	M = 6010
	N = 110
)
const INF = 0x3f3f3f3f

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func networkDelayTime(times [][]int, n int, k int) int {
	var head [N]int
	var e [M]int
	var next [M]int
	var weight [M]int
	var idx int
	add := func(a, b, w int) {
		e[idx] = b
		// 边
		next[idx] = head[a]
		// 节点
		head[a] = idx
		weight[idx] = w
		idx++
	}
	dist := [N]int{}
	for i := 1; i <= n; i++ {
		head[i] = -1
	}
	for _, time := range times {
		add(time[0], time[1], time[2])
	}
	// SPFA
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	vis := make([]bool, N)
	q := make([]int, 0)
	q = append(q, k)
	vis[k] = true
	dist[k] = 0
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		vis[node] = false
		for i := head[node]; i != -1; i = next[i] {
			j := e[i]
			if dist[j] > dist[node]+weight[i] {
				dist[j] = dist[node] + weight[i]
				if vis[j] {
					continue
				}
				q = append(q, j)
				vis[j] = true
			}
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

func main() {
	networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)
}
