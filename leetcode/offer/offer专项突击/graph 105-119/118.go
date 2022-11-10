package graph_105_119

// 并查集
// 其实是找第一条让其成环的边
func findRedundantConnection2(edges [][]int) []int {
	parent := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x1, x2 int) bool {
		x, y := find(x1), find(x2)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0]-1, e[1]-1) {
			return e
		}
	}
	return []int{}
}

// dfs
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	graph := make([][]int, n)
	for f := range graph {
		graph[f] = make([]int, 0)
	}
	for _, e := range edges {
		v1 := e[0] - 1
		v2 := e[1] - 1
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
		vis := make([]int, n)
		parent := make([]int, n)
		for i := 0; i < n; i++ {
			parent[i] = -1
		}
		has := false

		var dfs func(node int)
		dfs = func(node int) {
			vis[node] = 1
			for _, v := range graph[node] {
				if v == node {
					continue
				}
				if vis[v] == 1 && parent[node] != v {
					has = true
					return
				} else if vis[v] == 0 {
					parent[v] = node
					dfs(v)
				}
			}
			vis[node] = 2
		}
		dfs(v1)

		if has {
			return []int{e[0], e[1]}
		}
		for i := 0; i < n; i++ {
			vis[i] = 0
		}
		for i := 0; i < n; i++ {
			parent[i] = -1
		}

	}

	return []int{}
}
