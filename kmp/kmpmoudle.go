package kmp

func Kmp(next []int, pattern, str string) int {
	i, j := 0, 0
	for j < len(pattern) {
		if str[i] == pattern[j] {
			i += 1
			j += 1
		} else if j > 0 {
			j = next[j-1]
		} else {
			i += 1
		}
	}
	if j == len(pattern) {
		return i - j
	}
	return -1
}

func GetNext(pattern string) []int {
	next := make([]int, 0)
	next = append(next, 0)
	i := 1
	prefixLen := 0
	for i < len(pattern) {
		if pattern[prefixLen] == pattern[i] {
			prefixLen += 1
			i += 1
			next = append(next, prefixLen)
		} else {
			if prefixLen == 0 {
				next = append(next, 0)
				i += 1
			} else {
				prefixLen = next[prefixLen-1]
			}
		}
	}
	return next
}
