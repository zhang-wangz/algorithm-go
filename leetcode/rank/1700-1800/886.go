package main

// dfs 也可以并查集， 对不喜欢的反个进行union，如果发现当前不喜欢的a和b在同一个集合则返回false
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n+1)
	for _, e := range dislikes {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	color := map[int]int{}
	var dfs func(node, c int) bool
	dfs = func(node, c int) bool {
		if v, ok := color[node]; ok {
			return v == c
		}
		color[node] = c
		for _, n := range g[node] {
			if !dfs(n, c^1) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if _, ok := color[i]; !ok && !dfs(i, 0) {
			return false
		}
	}
	return true
}
