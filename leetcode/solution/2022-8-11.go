package main

import "unicode"

// https://leetcode.cn/problems/reformat-the-string/
func reformat(s string) string {
	sumDigit := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			sumDigit++
		}
	}
	sumAl := len(s) - sumDigit
	if abs(sumAl-sumDigit) > 1 {
		return ""
	}
	flag := sumDigit > sumAl
	t := []byte(s)
	for i, j := 0, 1; i < len(t); i += 2 {
		// 如果数字多，且当前i不是数字
		if unicode.IsDigit(rune(t[i])) != flag {
			// 数字多，且当前j不是数字，就+=2，直到j是数字
			for unicode.IsDigit(rune(t[j])) != flag {
				j += 2
			}
			// 交换数字和字母
			t[i], t[j] = t[j], t[i]
		}
	}
	return string(t)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
