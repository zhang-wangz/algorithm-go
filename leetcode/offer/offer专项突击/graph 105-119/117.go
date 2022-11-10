package graph_105_119

func numSimilarGroups(strs []string) int {
	parent := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x1, x2 int) {
		parent[find(x1)] = find(x2)
	}
	as := 0
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if isValid(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}
	for i := 0; i < len(parent); i++ {
		if parent[i] == i {
			as++
		}
	}
	return as
}

func isValid(w1 string, w2 string) bool {
	var f1, f2 byte
	f1 = '0'
	f2 = '0'
	if len(w1) != len(w2) {
		return false
	}
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] && f1 == '0' && f2 == '0' {
			f1 = w1[i]
			f2 = w2[i]
		} else if w1[i] != w2[i] {
			if w1[i] != f2 || w2[i] != f1 {
				return false
			}
		}
	}
	return true
}
