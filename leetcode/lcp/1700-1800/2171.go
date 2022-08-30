package main

import "sort"

// 2171
// 获取能保留的最多的豆，减一下就是取最少的豆
func minimumRemoval(beans []int) int64 {
	sum := 0
	mi := 0
	sort.Ints(beans)
	for i, v := range beans {
		sum += v
		bean := (len(beans) - i) * v
		if bean > mi {
			mi = bean
		}
	}
	return int64(sum - mi)
}
