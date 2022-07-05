package dfs_79_87

func combine(n int, k int) (ans [][]int) {
	path := make([]int, 0)
	var dfs func(n int, startIdx int)
	dfs = func(n int, start int) {
		if len(path)+(n-start+1) < k {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(n, i+1)
			path = path[:len(path)-1]
		}
		return
	}
	dfs(n, 1)
	return
}
