package _022_07

//https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/

func oddCells(m int, n int, indices [][]int) int {
	r := map[int]int{}
	c := map[int]int{}
	for _, v := range indices {
		r[v[0]]++
		c[v[1]]++
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (r[i]+c[j])&1 == 1 {
				ans++
			}
		}
	}
	return ans
}
