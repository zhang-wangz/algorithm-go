package contest

// https://leetcode.cn/problems/match-substring-after-replacement/
// 给你两个字符串 s 和 sub 。同时给你一个二维字符数组 mappings ，
// 其中 mappings[i] = [oldi, newi] 表示你可以替换 sub 中任意数目的 oldi 个字符，替换成 newi 。sub 中每个字符 不能 被替换超过一次。
// 如果使用 mappings 替换 0 个或者若干个字符，可以将 sub 变成 s 的一个子字符串，请你返回 true，否则返回 false 。
// 一个 子字符串 是字符串中连续非空的字符序列。

// kmp
func MatchReplacement(s string, sub string, mappings [][]byte) bool {
	m := len(s)
	n := len(sub)
	ma := map[byte][]byte{}
	for _, e := range mappings {
		ma[e[0]] = append(ma[e[0]], e[1])
	}
	for i := 0; i+n-1 < m; i++ {
		if isValid(s[i:i+n], sub, ma) {
			return true
		}
	}
	return false
}

func isValid(a, b string, m map[byte][]byte) bool {
	// 字符是否被替换
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			// 可以替换
			if len(m[b[i]]) != 0 {
				var f bool
				for j := 0; j < len(m[b[i]]); j++ {
					if m[b[i]][j] == a[i] {
						f = true
						break
					}
				}
				if !f {
					return false
				}
			}
			if _, ok := m[b[i]]; !ok {
				return false
			}
		}
	}
	return true
}
