package main

func isBipartite2(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			q := make([]int, 0)
			q = append(q, i)
			color[i] = 1
			for len(q) != 0 {
				node := q[0]
				q = q[1:]
				dc := 1
				if color[node] == 1 {
					dc = 2
				}
				for _, nei := range graph[node] {
					if color[nei] == 0 {
						color[nei] = dc
						q = append(q, nei)
					} else if color[nei] != dc {
						return false
					}
				}
			}
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	valid := true
	var dfs func(node int, c int)
	dfs = func(node int, c int) {
		color[node] = c
		dc := 1
		if c == 1 {
			dc = 2
		}
		for i := 0; i < len(graph[node]); i++ {
			nei := graph[node][i]
			if color[nei] == 0 {
				dfs(nei, dc)
				if !valid {
					return
				}
			} else if color[nei] != dc {
				valid = false
				return
			}
		}
	}
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			dfs(i, 1)
		}
	}
	return valid
}
