package dfs_79_87

func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	var dfs func(int)
	path := make([]string, 0)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
