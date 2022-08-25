package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) (ans []int) {
	red, blue := make([][]int, n), make([][]int, n)
	for _, e := range redEdges {
		red[e[0]] = append(red[e[0]], e[1])
	}
	for _, e := range blueEdges {
		blue[e[0]] = append(blue[e[0]], e[1])
	}
	// 两种状态下的dis距离, 0到该点的距离
	dis := [2][]int{}
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}
	// node, dis, color
	dis[0][0] = 0
	dis[1][0] = 0
	q := [][]int{{0, 0, 0}, {0, 0, 1}}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		color := node[2]
		// 如果该点是红色，去找蓝色
		if color == 0 {
			for _, e := range blue[node[0]] {
				if dis[1][e] == -1 {
					dis[1][e] = node[1] + 1
					q = append(q, []int{e, node[1] + 1, 1})
				}
			}
		} else {
			for _, e := range red[node[0]] {
				if dis[0][e] == -1 {
					dis[0][e] = node[1] + 1
					q = append(q, []int{e, node[1] + 1, 0})
				}
			}
		}
	}
	for i := range dis[0] {
		if dis[0][i] == -1 || dis[1][i] == -1 {
			ans = append(ans, max(dis[0][i], dis[1][i]))
		} else {
			ans = append(ans, min(dis[0][i], dis[1][i]))
		}
	}
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
