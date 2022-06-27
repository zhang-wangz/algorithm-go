package offer

// https://leetcode.cn/problems/a7VOhD/
// 暴力枚举
func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			a := s[j:i]
			if validPalindrome(a) {
				cnt++
			}
		}
	}
	return cnt
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// 中心拓展
func countSubstrings2(s string) int {
	cnt := 0
	for i := 0; i <= 2*len(s)-2; i++ {
		l := i / 2
		r := l + i%2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			cnt++
		}
	}
	return cnt
}
