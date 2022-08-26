package main

// 764
// 遍历两遍，记数统计
func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	vis := map[int]bool{}
	for _, e := range mines {
		vis[e[0]*n+e[1]] = true
	}
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if !vis[i*n+j] {
				cnt++
				g[i][j] = cnt
			} else {
				cnt = 0
				g[i][j] = 0
			}
		}
		cnt = 0
		for j := n - 1; j >= 0; j-- {
			if !vis[i*n+j] {
				cnt++
				g[i][j] = min(g[i][j], cnt)
			} else {
				cnt = 0
				g[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		cnt := 0
		for i := 0; i < n; i++ {
			if !vis[i*n+j] {
				cnt++
				g[i][j] = min(g[i][j], cnt)
			} else {
				cnt = 0
				g[i][j] = 0
			}
		}
		cnt = 0
		for i := n - 1; i >= 0; i-- {
			if !vis[i*n+j] {
				cnt++
				g[i][j] = min(g[i][j], cnt)
			} else {
				cnt = 0
				g[i][j] = 0
			}
			ans = max(g[i][j], ans)
		}
	}
	return
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
