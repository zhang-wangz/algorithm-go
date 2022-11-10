package main

import "math"

// 842
func splitIntoFibonacci(num string) (ans []int) {
	n := len(num)
	var dfs func(idx, sum, pre int) bool
	dfs = func(idx, sum, pre int) bool {
		if idx == n {
			return len(ans) >= 3
		}
		cur := 0
		for i := idx; i < n; i++ {
			if i > idx && num[idx] == '0' {
				break
			}
			cur = cur*10 + int(num[i]-'0')
			if cur > math.MaxInt32 {
				break
			}
			if len(ans) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}
			ans = append(ans, cur)
			if dfs(i+1, pre+cur, cur) {
				return true
			}
			ans = ans[:len(ans)-1]
		}
		return false
	}
	dfs(0, 0, 0)
	return
}
