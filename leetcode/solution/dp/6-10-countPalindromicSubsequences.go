package dp

// https://leetcode.cn/problems/count-different-palindromic-subsequences/

var m = map[string]bool{}

func CountPalindromicSubsequences(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+5)
	for i := 1; i <= len(s); i++ {
		cnt := dp[i-1]
		for j := 0; j < i; j++ {
			if v, ok := m[s[:j]+string(s[i-1])]; ok {
				if v == true {
					continue
				}
			}
			if isValid(s[:j] + string(s[i-1])) {
				cnt++
			}
		}
		dp[i] = cnt
	}
	return dp[len(s)]
}

func isValid(s string) bool {
	if v, ok := m[s]; ok {
		return v
	}
	i, j := 0, len(s)
	for i < j {
		if s[i] != s[j] {
			m[s] = false
			return false
		}
		i++
		j++
	}
	m[s] = true
	return true
}
