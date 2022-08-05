package tusen2

func findLUSlength(strs []string) int {
	isSub := func(s, t string) bool {
		if len(s) > len(t) {
			return false
		}
		sidx := 0
		for tidx := range t {
			if s[sidx] == t[tidx] {
				if sidx++; sidx == len(s) {
					return true
				}
			}
		}
		return false
	}
	n := len(strs)
	ans := -1
out:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isSub(strs[i], strs[j]) {
				continue out
			}
		}
		if len(strs[i]) > ans {
			ans = len(strs[i])
		}
	}
	return ans
}
