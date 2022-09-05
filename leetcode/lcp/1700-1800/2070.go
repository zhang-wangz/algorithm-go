package main

import "sort"

// 2070 每一个查询的最大美丽值
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	for i, q := range queries {
		queries[i] = q<<32 | i
	}
	sort.Ints(queries)
	mx, i := 0, 0
	ans := make([]int, len(queries))
	for _, q := range queries {
		for ; i < len(items) && items[i][0] <= q>>32; i++ {
			if items[i][1] > mx {
				mx = items[i][1]
			}
		}
		ans[q&(1<<32-1)] = mx
	}
	return ans
}
