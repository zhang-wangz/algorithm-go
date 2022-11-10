package _022_05

// 由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
//
// 如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
// 如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
// 给定一个字符串 s ，重构排列 perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。

func diStringMatch(s string) []int {
	if s == " " {
		return []int{}
	}
	res := make([]int, 0)
	n := len(s)
	left := 0
	right := n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			res = append(res, left)
			left++
		} else {
			res = append(res, right)
			right--
		}
	}
	res = append(res, left)
	return res
}
