package main

import "strings"

func wordPattern(pattern string, s string) bool {
	b := strings.Split(s, " ")
	if len(b) != len(pattern) {
		return false
	}
	m1, m2 := map[string]string{}, map[string]string{}
	for i, s1 := range b {
		s2 := string(pattern[i])
		if v1, ok1 := m1[s1]; ok1 && v1 != s2 {
			return false
		}
		if v2, ok2 := m2[s2]; ok2 && v2 != s1 {
			return false
		}
		m1[s1] = s2
		m2[s2] = s1
	}
	return true
}
