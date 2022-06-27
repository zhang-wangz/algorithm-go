package _022_6_10

// https://leetcode.cn/problems/match-substring-after-replacement/
// 给你两个字符串 s 和 sub 。同时给你一个二维字符数组 mappings ，
// 其中 mappings[i] = [oldi, newi] 表示你可以替换 sub 中任意数目的 oldi 个字符，替换成 newi 。sub 中每个字符 不能 被替换超过一次。
// 如果使用 mappings 替换 0 个或者若干个字符，可以将 sub 变成 s 的一个子字符串，请你返回 true，否则返回 false 。
// 一个 子字符串 是字符串中连续非空的字符序列。

// 枚举
func MatchReplacement(s string, sub string, mappings [][]byte) bool {
	mp := ['z' + 1]['z' + 1]bool{}
	for _, e := range mappings {
		mp[e[0]][e[1]] = true
	}
next:
	for i := len(sub); i <= len(s); i++ {
		start := i - len(sub)
		for idx, p := range s[start:i] {
			if byte(p) != sub[idx] && !mp[sub[idx]][p] {
				continue next
			}
		}
		// 到这里说明匹配成功了
		return true
	}
	// 还没返回成功说明没有成功的
	return false
}
