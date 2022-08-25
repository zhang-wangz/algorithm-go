package main

import (
	"strings"
)

func maxScore(s string) (ans int) {
	score := 1 - int(s[0]) + strings.Count(s[1:], "1")
	for _, c := range s[1 : len(s)-1] {
		if c == '0' {
			score++
		} else {
			score--
		}
		ans = max(ans, score)
	}
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
