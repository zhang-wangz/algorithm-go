package main

import "math"

// 1781 所有子字符串美丽值之和
func beautySum(s string) (ans int) {
	// 暴力每一个子串
	n := len(s)
	for le := 3; le <= n; le++ {
		l, mi, ma := 0, math.MaxInt, math.MinInt
		cnt := [26]int{}
		for r := range s {
			cnt[s[r]-'a']++
			if r-l+1 == le {
				mi, ma = math.MaxInt, math.MaxInt
				for i := 0; i < 26; i++ {
					if cnt[i] > 0 {
						mi = min(mi, cnt[i])
						ma = max(ma, cnt[i])
					}
				}
				ans += ma - mi
				cnt[s[l]-'a']--
				l++
			}
		}
	}
	return
}
