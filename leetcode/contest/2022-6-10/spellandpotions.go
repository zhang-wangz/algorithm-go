package _022_6_10

import "sort"

func spellsAndpotions(spells []int, potions []int, success int64) []int {
	res := make([]int, 0)
	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		l := 0
		r := len(potions)
		// 找 >= success的最小下标
		for l < r {
			mid := l + (r-l)/2
			if potions[mid]*spells[i] >= int(success) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == len(potions) {
			res = append(res, 0)
		} else {
			res = append(res, len(potions)-l)
		}
		// 或者
		//for l+1 < r {
		//	mid := l + (r-l)/2
		//	if potions[mid]*spells[i] >= int(success) {
		//		r = mid
		//	} else {
		//		l = mid
		//	}
		//}
		//if potions[l]*spells[i] >= int(success) {
		//	res = append(res, len(potions)-l)
		//} else if potions[r]*spells[i] >= int(success) {
		//	res = append(res, len(potions)-r)
		//}else {
		//	res = append(res, 0)
		//}

	}
	return res
}
