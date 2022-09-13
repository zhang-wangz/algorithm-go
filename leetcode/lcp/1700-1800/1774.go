package main

import (
	"math"
	"sort"
)

// 1774
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	n := len(baseCosts)
	m := len(toppingCosts)
	sort.Ints(toppingCosts)
	sort.Ints(baseCosts)

	ret := baseCosts[0]
	// 每个基料对应最多3^m的方案
	mx := int(math.Pow(3, float64(m)))
	for i := 0; i < n; i++ {
		s := 0
		// 遍历所有辅料的方案
		var cr int
		for s < mx {
			cs := s
			cr = baseCosts[i]
			for j := 0; j < m; j++ {
				c := cs % 3
				cs /= 3
				cr += c * toppingCosts[j]
			}
			if math.Abs(float64(cr-target)) < math.Abs(float64(ret-target)) {
				ret = cr
			} else if math.Abs(float64(cr-target)) == math.Abs(float64(ret-target)) {
				if cr < ret {
					ret = cr
				}
			}
			if ret == target {
				return target
			}
			s++
		}
	}
	return ret
}
