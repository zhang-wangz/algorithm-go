package main

// https://leetcode.cn/problems/cyJERH/

func minFlipsMonoIncr(s string) int {
	n := len(s)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		a1 := a + int(s[i]-'0')
		b1 := min(a, b) + (int(s[i]-'0') ^ 1)
		a, b = a1, b1
	}
	return min(a, b)
}
