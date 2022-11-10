package main

// 根据余数进行匹配
func canArrange(arr []int, k int) bool {
	mod := make([]int, k)
	for _, n := range arr {
		mod[(n%k+k)%k]++
	}
	for i := 1; i <= k/2; i++ {
		if mod[i] != mod[k-i] {
			return false
		}
	}
	return mod[0]%2 == 0
}

// 超时
func canArrange2(arr []int, k int) bool {
	// 二分图匹配
	n := len(arr)
	vis := make([]bool, n)
	match := make([]int, n)
	for i := range match {
		match[i] = -1
	}
	var dfs func(u int) bool
	dfs = func(u int) bool {
		for i := range arr {
			if !vis[i] && (arr[i]+arr[u])%k == 0 {
				vis[i] = true
				if match[i] == -1 || dfs(match[i]) {
					match[i] = u
					return true
				}
			}
		}
		return false
	}
	for i := range arr {
		if !vis[i] && !dfs(i) {
			return false
		}
	}
	return true
}
