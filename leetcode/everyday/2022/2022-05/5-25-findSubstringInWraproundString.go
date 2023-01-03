package _022_05

// https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
// 把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的：
// "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." .
// 现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。

// dp
// 维护以每个小写字母为结尾的dp，表示以该字母结尾的有多少个子窜
func findSubstringInWraproundString2(p string) int {
	if len(p) == 0 {
		return 0
	}
	dp := make([]int, 26)
	dp[p[0]-'a'] = 1
	cur := 1
	for i := 1; i < len(p); i++ {
		if (p[i]-p[i-1]+25)%26 == 0 {
			cur++
		} else {
			cur = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], cur)
	}
	sum := 0
	for _, v := range dp {
		sum += v
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// TLE
func findSubstringInWraproundString(p string) int {
	vis := map[string]bool{}
	cnt := 0
	for i := 0; i < len(p); i++ {
		for j := i + 1; j <= len(p); j++ {
			if isValid(p[i:j]) {
				if _, ok := vis[p[i:j]]; !ok {
					vis[p[i:j]] = true
					cnt++
				}
			}
		}
	}
	return cnt
}

func isValid(p string) bool {
	if len(p) == 0 {
		return false
	}
	for i := 0; i < len(p)-1; i++ {
		if p[i] == 'z' && p[i+1] == 'a' {
			continue
		}
		if p[i+1]-p[i] != 1 {
			return false
		}
	}
	return true
}
