package main

import "strings"

func replaceSpace(s string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
