package _022_08

import "sort"

func orderlyQueue(s string, k int) string {
	b := []byte(s)
	if k == 1 {
		for i := 0; i < len(b); i++ {
			if b = append(b[1:], b[0]); string(b) < s {
				s = string(b)
			}
		}
		return s
	} else {
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		return string(b)
	}
}
