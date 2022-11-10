package main

import "sort"

func bagOfTokensScore(tokens []int, power int) (ans int) {
	sort.Ints(tokens)
	if len(tokens) == 0 {
		return 0
	}
	if len(tokens) == 1 {
		if power > tokens[0] {
			return 1
		}
	}
	fs := 0
	l, r := 0, len(tokens)-1
	for l <= r {
		t1, t2 := tokens[l], tokens[r]
		if power >= t1 {
			power -= t1
			fs += 1
			ans = max(ans, fs)
			l++
		} else if fs > 0 {
			power += t2
			fs -= 1
			r--
		} else {
			break
		}
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
