package solution

// https://leetcode.cn/problems/find-and-replace-pattern/

func FindAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	m := map[byte]int{}
	// 0表示A，1表示B，以此类推
	cnt := 0
	arr := make([]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		if _, ok := m[pattern[i]]; !ok {
			m[pattern[i]] = cnt
			cnt++
		}
		arr[i] = m[pattern[i]]
	}
	for i := 0; i < len(words); i++ {
		tmp := map[byte]int{}
		cnt := 0
		arrtmp := make([]int, len(words[i]))
		var j = 0
		for ; j < len(words[i]); j++ {
			if _, ok := tmp[words[i][j]]; !ok {
				tmp[words[i][j]] = cnt
				cnt++
			}
			arrtmp[j] = tmp[words[i][j]]
			if arrtmp[j] != arr[j] {
				break
			}
		}
		if j == len(words[i]) {
			res = append(res, words[i])
		}
	}
	return res
}
