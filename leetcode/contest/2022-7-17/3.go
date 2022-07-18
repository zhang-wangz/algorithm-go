package _022_7_17

import (
	"sort"
)

type pair struct {
	v   string
	idx int
}

func smallestTrimmedNumbers(nums []string, queries [][]int) (ans []int) {
	for i := 0; i < len(queries); i++ {
		tmp := make([]pair, len(nums))
		l := queries[i][1]
		k := queries[i][0]
		for f := range nums {
			v := nums[f][len(nums[f])-l : len(nums[f])]
			p := pair{v: v, idx: f}
			tmp[f] = p
		}
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i].v < tmp[j].v {
				return tmp[i].v < tmp[j].v
			} else if tmp[i].v == tmp[j].v {
				return tmp[i].idx < tmp[j].idx
			} else {
				return tmp[i].v < tmp[j].v
			}
		})
		ans = append(ans, tmp[k-1].idx)
	}
	return
}
