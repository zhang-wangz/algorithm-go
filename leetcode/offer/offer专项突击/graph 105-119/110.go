package graph_105_119

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	path := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for _, y := range graph[x] {
			path = append(path, y)
			dfs(y)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
