package main

// 1914 循环轮转
// 给你一个大小为 m x n 的整数矩阵 grid ，其中 m 和 n 都是 偶数 ；另给你一个整数 k 。
func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	layer := min(m, n) / 2
	for la := 0; la < layer; la++ {
		r, c, val := []int{}, []int{}, []int{}
		for i := la; i < n-la; i++ {
			r = append(r, la)
			c = append(c, i)
			val = append(val, grid[la][i])
		}
		for i := la + 1; i < m-la; i++ {
			r = append(r, i)
			c = append(c, n-la-1)
			val = append(val, grid[i][n-la-1])
		}
		for i := n - la - 2; i >= la; i-- {
			r = append(r, m-la-1)
			c = append(c, i)
			val = append(val, grid[m-la-1][i])
		}
		for i := m - la - 2; i > la; i-- {
			r = append(r, i)
			c = append(c, la)
			val = append(val, grid[i][la])
		}
		tot := len(val)
		kk := k % tot
		for i := 0; i < tot; i++ {
			idx := (i + kk) % tot
			grid[r[i]][c[i]] = val[idx]
		}
	}
	return grid
}

//func main() {
//	rotateGrid([][]int{{40,10},{30,20}}, 1)
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
