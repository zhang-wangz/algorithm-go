package main

import (
	"bytes"
)

// 2075 tle
func decodeCiphertext(encodedText string, rows int) string {
	m := rows
	n := len(encodedText) / m
	ans := []byte{}
	for i, j, k := 0, 0, 0; k < n; {
		ans = append(ans, encodedText[i*n+j])
		i++
		j++
		if i == m || j == n {
			k++
			i, j = 0, k
		}
	}
	return string(bytes.TrimRight(ans, " "))
}
