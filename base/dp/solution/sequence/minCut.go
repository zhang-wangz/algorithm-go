package sequence

// https://leetcode.cn/problems/palindrome-partitioning-ii/
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
// 返回符合要求的 最少分割次数 。
func MinCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// 设置最大切割数
		r[i] = i
		for j := 0; j <= i; j++ {
			if isPalindrome(s, j, i) {
				if j == 0 {
					// 如果是0开始的，和最前面的0开始比较
					r[i] = min(r[i], 0)
				} else {
					// 不是0的话，需要回文前的最少+1进行比较
					r[i] = min(r[i], r[j-1]+1)
				}
			}
		}
	}
	return r[len(s)-1]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
