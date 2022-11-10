package main

// 1653
// 使得编程aaaabbb形式
func minimumDeletions(s string) int {
	// 就是使得编程01形式的最少步骤
	n := len(s)
	dp0 := int(s[0] - 'a')
	dp1 := 0
	if s[0] == 'a' {
		dp1 = 1
	}
	for i := 1; i < n; i++ {
		t := dp0
		if s[i] == 'a' {
			dp0 = dp0
			dp1 = min(t, dp1) + 1
		} else {
			dp0 = dp0 + 1
			dp1 = min(t, dp1)
		}
	}
	return min(dp0, dp1)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
