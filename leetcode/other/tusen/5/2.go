package main

// https://leetcode.cn/problems/rotating-the-box/

func rotateTheBox(a [][]byte) [][]byte {
	n, m := len(a), len(a[0])
	ans := make([][]byte, m)
	for i := range ans {
		ans[i] = make([]byte, n)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
	}
	for i, r := range a {
		for j := 0; j < m; j++ {
			c := 0
			// 注意这里 j 和外层循环的 j 是同一个变量，因此时间复杂度为 O(nm)
			for ; j < m && r[j] != '*'; j++ {
				if r[j] == '#' {
					c++
				}
			}
			if j < m {
				ans[j][n-1-i] = '*'
			}
			for k := j - 1; c > 0; k-- {
				ans[k][n-1-i] = '#'
				c--
			}
		}
	}
	return ans
}
