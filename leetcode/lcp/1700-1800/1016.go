package main

import (
	"strconv"
	"strings"
)

// dfs每个子串表示的数，用一个mask（1000位超过int范围了 ）记录，如果mask等于2^n-1则true
func queryString2(s string, n int) bool {
	for i := 1; i <= n; i++ {
		a := strconv.FormatInt(int64(i), 2)
		if !strings.Contains(s, a) {
			return false
		}
	}
	return true
}
