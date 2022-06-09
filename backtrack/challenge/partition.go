package challenge

// https://leetcode.cn/problems/palindrome-partitioning/
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串 。返回 s 所有可能的分割方案。
// 回文串 是正着读和反着读都一样的字符串。

func Partition(s string) [][]string {
	res := make([][]string, 0)
	list := make([]string, 0)
	dfs5(s, &res, list)
	return res
}

func dfs5(s string, res *[][]string, list []string) {
	if s == "" {
		ans := make([]string, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}
	for i := 0; i < len(s); i++ {
		str := s[:i+1]
		if isValid(str) {
			list = append(list, str)
			dfs5(s[i+1:], res, list)
			list = list[:len(list)-1]
		}
	}
}

func isValid(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
