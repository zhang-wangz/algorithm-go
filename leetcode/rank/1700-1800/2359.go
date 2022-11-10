package main

// 2359 找到离给定两个节点最近的节点
func closestMeetingNode(edges []int, node1 int, node2 int) (ans int) {
	const inf = 0x3f3f3f3f
	bfs := func(node int) []int {
		dis := make([]int, len(edges))
		for i := range dis {
			dis[i] = inf
		}
		d := 0
		for node != -1 && dis[node] == inf {
			dis[node] = d
			d += 1
			node = edges[node]
		}
		return dis
	}
	ans = -1
	n1 := bfs(node1)
	n2 := bfs(node2)
	md := len(edges)
	for i := range n1 {
		mx := max(n1[i], n2[i])
		if mx < md {
			md = mx
			ans = i
		}
	}
	return
}
