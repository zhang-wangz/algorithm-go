package contest

import "sort"

func spellsAndpotions(spells []int, potions []int, success int64) []int {
	res := make([]int, 0)
	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		l := 0
		r := len(potions) - 1
		// 找第一个 >= success的值
		for l+1 < r {
			mid := l + (r-l)/2
			if potions[mid]*spells[i] >= int(success) {
				r = mid
			} else {
				l = mid
			}
		}
		if potions[l]*spells[i] >= int(success) {
			res = append(res, l)
		} else if potions[r]*spells[i] >= int(success) {
			res = append(res, r)
		}

	}
	return res
}
