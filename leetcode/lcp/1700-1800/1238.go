package main

import "strconv"

// 其实就是求格雷码，格雷码的特性是首个和最后一个2进制差一个1并且邻位也是2进制各自差个1
// 最简单就是直接套公式
func circularPermutation2(n int, start int) (ans []int) {
	for i := 0; i < 1<<n; i++ {
		ans = append(ans, i^(i>>1)^start)
	}
	return
}

// 递归
func circularPermutation(n int, start int) (ans []int) {
	var dfs func(n int) []string
	dfs = func(n int) []string {
		grey := make([]string, 1<<n)
		if n == 1 {
			grey[0] = "0"
			grey[1] = "1"
			return grey
		}
		last := dfs(n - 1)
		for i := 0; i < len(last); i++ {
			grey[i] = "0" + last[i]
			grey[len(grey)-1-i] = "1" + last[i]
		}
		return grey
	}
	grey := dfs(n)
	for i := 0; i < len(grey); i++ {
		n, _ := strconv.ParseInt(grey[i], 2, 0)
		ans = append(ans, int(n)^start)
	}
	return
}
