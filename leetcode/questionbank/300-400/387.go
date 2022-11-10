package main

func firstUniqChar(s string) int {
	cti := map[int]int{}
	for _, c := range s {
		cti[int(c-'a')]++
	}
	for i, c := range s {
		if cti[int(c-'a')] == 1 {
			return i
		}
	}
	return -1
}
