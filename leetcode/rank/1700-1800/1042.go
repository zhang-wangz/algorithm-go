package main

// 1042 不邻接植花
func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, p := range paths {
		a, b := p[0]-1, p[1]-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		color := [5]int{0}
		for _, t := range g[i] {
			color[ans[t]] = 1
		}
		for c := 1; c <= 4; c++ {
			if color[c] == 0 {
				ans[i] = c
				break
			}
		}
	}
	return ans
}
