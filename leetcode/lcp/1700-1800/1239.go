package main

import "math/bits"

// 1239 串联字符串的最大长度
func maxLength(arr []string) (ans int) {
	masks := []int{0}
out:
	for _, s := range arr {
		mask := 0
		for _, c := range s {
			ch := int(c - 'a')
			if mask>>ch&1 == 1 {
				continue out
			}
			mask |= 1 << ch
		}
		for _, m := range masks {
			if m&mask == 0 {
				masks = append(masks, m|mask)
				ans = max(ans, bits.OnesCount(uint(m|mask)))
			}
		}
	}
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
