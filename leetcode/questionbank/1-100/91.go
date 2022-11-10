package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	// 设置dp[n]有多少种解码方式
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i, c := range s {
		if c != '0' {
			f[i+1] += f[i]
		}
		if i > 0 && s[i-1] != '0' && int(s[i-1]-'0')*10+int(c-'0') <= 26 {
			f[i+1] += f[i-1]
		}
	}
	return f[n]
}
