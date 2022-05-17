package solution

// https://leetcode.cn/problems/verifying-an-alien-dictionary/
// 某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
//
// 给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；
// 否则，返回 false。

func isAlienSorted(words []string, order string) bool {
	m := make(map[byte]int, len(order))
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j == len(words[i+1]) {
				return false
			}
			a := m[words[i][j]]
			b := m[words[i+1][j]]
			if a > b {
				return false
			} else if a < b {
				break
			}
		}
	}
	return true
}
