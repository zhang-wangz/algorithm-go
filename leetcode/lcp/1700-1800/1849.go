package main

import (
	"strconv"
	"strings"
)

// 1849
// 枚举第一个值
func splitString(s string) bool {
next:
	for i := range s {
		v, _ := strconv.Atoi(s[:i])
		v--
		for t := s[i:]; t != ""; v-- {
			for len(t) > 1 && t[0] == '0' {
				t = t[1:]
			}
			tmp := strconv.Itoa(v)
			if !strings.HasPrefix(t, tmp) {
				continue next
			}
			t = t[len(tmp):]
		}
		return true
	}
	return false
}
