package main

// 1391
// 网络中是否存在有效路径
func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	parents := make([]int, m*n)
	for i, _ := range parents {
		parents[i] = i
	}
	var find func(a int) int

	find = func(a int) int {
		if a != parents[a] {
			parents[a] = find(parents[a])
		}
		return parents[a]
	}
	isUnion := func(a, b int) bool {
		return find(a) == find(b)
	}
	union := func(a, b int) {
		parents[find(a)] = find(b)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 当右边是通向的，右边是左向的，表面可达
			if grid[i][j] == 1 || grid[i][j] == 4 || grid[i][j] == 6 {

				if j+1 < n && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
					union(i*n+j, i*n+j+1)
				}
			}
			// 当下面是通向的，则下面是上面也是通的
			if grid[i][j] == 2 || grid[i][j] == 3 || grid[i][j] == 4 {
				if i+1 < m && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
					union((i+1)*n+j, i*n+j)
				}
			}
		}
	}

	return isUnion(0, n*m-1)
}
