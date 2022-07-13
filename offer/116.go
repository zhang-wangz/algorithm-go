package main

// 并查集
func findCircleNum3(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	ans := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == i {
			ans++
		}
	}
	return ans
}

// dfs
func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	vis := make([]bool, n)
	var dfs func(from int)
	dfs = func(from int) {
		q := []int{from}
		for len(q) != 0 {
			x := q[0]
			q = q[1:]
			vis[x] = true
			for i, v := range isConnected[x] {
				if !vis[i] && v == 1 {
					dfs(i)
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
			ans++
		}
	}
	return ans
}

// bfs
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	vis := make([]bool, n)
	bfs := func(node int) {
		q := []int{node}
		for len(q) != 0 {
			x := q[0]
			q = q[1:]
			vis[x] = true
			for i, v := range isConnected[x] {
				if !vis[i] && v == 1 {
					q = append(q, i)
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			bfs(i)
			ans++
		}
	}
	return ans
}
