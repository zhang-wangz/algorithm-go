package dfs_79_87

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) (ans []string) {
	if s == "" {
		return ans
	}
	if len(s) > 4*3 {
		return ans
	}
	path := make([]string, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(s)-idx > (4-len(path))*3 {
			return
		}
		if len(path) == 4 && idx == len(s) {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		for i := idx; i < len(s); i++ {
			if ok, n := isValid(s[idx : i+1]); ok {
				path = append(path, strconv.Itoa(n))
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func isValid(s string) (bool, int) {
	n := 0
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '0' && len(s) != 1 {
			return false, -1
		}
		n = n*10 + int(s[i]-'0')
	}
	if n >= 0 && n <= 255 {
		return true, n
	}
	return false, -1
}
