package main

// 1638 统计只差一个字符的子串数目
// 从s遍历到t，分别遍历每次找到的数目
func countSubstrings(s string, t string) (ans int) {
	m, n := len(s), len(t)
	for del := -(m - 1); del < n; del++ {
		i, j := 0, 0
		if del > 0 {
			j = del
		} else {
			i = -del
		}
		// f统计数目，g统计后缀长度
		f, g := 0, 0
		for ; i < m && j < n; i, j = i+1, j+1 {
			if s[i] == t[j] {
				// f 不变
				g++
			} else {
				f = g + 1
				g = 0
			}
			ans += f
		}
	}
	return
}
