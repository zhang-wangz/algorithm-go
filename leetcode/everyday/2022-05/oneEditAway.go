package _022_05

import "math"

// https://leetcode.cn/problems/one-away-lcci/
// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
// 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if math.Abs(float64(m-n)) > 1 {
		return false
	}
	i, j := 0, 0
	var flag bool
	for i < m && j < n {
		if first[i] == second[j] {
			i++
			j++
		} else if flag {
			return false
		} else {
			// m比n长，所以不一样的地方在m处
			if m > n {
				i++
			} else if m < n {
				j++
			} else {
				i++
				j++
			}
			// 第一次需要编辑改为true
			flag = true
		}
	}
	return true
}
