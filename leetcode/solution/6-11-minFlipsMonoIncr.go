package solution

import "math"

// https://leetcode.cn/problems/flip-string-to-monotone-increasing/
// 如果一个二进制字符串，是以一些 0（可能没有 0）后面跟着一些 1（也可能没有 1）的形式组成的，那么该字符串是 单调递增 的。
// 给你一个二进制字符串 s，你可以将任何 0 翻转为 1 或者将 1 翻转为 0 。
// 返回使 s 单调递增的最小翻转次数。

func MinFlipsMonoIncr(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= len(s); i++ {
		dp[i][0] = dp[i-1][0] + int(s[i-1]-'0')
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + (int(s[i-1]-'0') ^ 1)
	}
	return min(dp[len(s)][0], dp[len(s)][1])
}

// lcs 和 lis
func MinFlipsMonoIncr2(s string) int {
	Max := 0
	g := make([]int, len(s)+2)
	for i := 0; i < len(g); i++ {
		g[i] = len(s) + 100
	}
	for i := 0; i < len(s); i++ {
		// [1,i+1)
		t := int(s[i] - '0')
		l, r := 1, i+1
		for l < r {
			mid := l + (r-l)>>1
			if g[mid] <= t {
				l = mid + 1
			} else {
				r = mid
			}
		}
		// l-1是小于等于t的最大len，len+1是当前元素的LIS
		le := l - 1 + 1
		// 此时g[le] > t
		// +1之后是最小下标
		//g[le] = min(g[le],t)，可以简写成
		g[le] = t
		Max = max(Max, le)
	}
	return len(s) - Max
}

// 前缀和
func MinFlipsMonoIncr3(s string) int {
	sum := make([]int, len(s)+1)
	ans := math.MaxInt
	for i := 1; i <= len(s); i++ {
		sum[i] = sum[i-1] + int(s[i-1]-'0')
	}
	for i := 1; i <= len(s); i++ {
		// 左边有多少1 要改成0
		l := sum[i-1]
		r := (len(s) - i) - (sum[len(s)] - sum[i])
		ans = min(ans, l+r)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
