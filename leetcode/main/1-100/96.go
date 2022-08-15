package main

// G和F进行归纳求解
// https://leetcode.cn/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i; j++ {
			g[i] = g[j-1] * g[i-j]
		}
	}
	return g[n]
}
