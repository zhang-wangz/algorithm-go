package main

// 1593
// 给你一个字符串 s ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。
// 字符串 s 拆分后可以得到若干非空子字符串，这些子字符串连接后应当能够还原为原字符串。
// 但是拆分出来的每个子字符串都必须是唯一的 。
func maxUniqueSplit(s string) (ans int) {
	n := len(s)
	s2b := map[string]struct{}{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			if len(s2b) > ans {
				ans = len(s2b)
			}
		}
		if len(s2b)+n-i < ans {
			return
		}
		for k := i; k < n; k++ {
			sub := s[i : k+1]
			if _, ok := s2b[sub]; !ok {
				s2b[sub] = struct{}{}
				dfs(i + 1)
				delete(s2b, sub)
			}
		}
	}
	dfs(0)
	return
}
