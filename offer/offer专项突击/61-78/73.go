package _1_78

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000001
	for l < r {
		mid := l + (r-l)>>1
		if eatBan(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 求符合条件的最小速度
func eatBan(piles []int, k int) int {
	h := 0
	for i := 0; i < len(piles); i++ {
		del := piles[i] / k
		if del == 0 {
			h += 1
		} else {
			h += del
		}
		if del != 0 && piles[i]%k != 0 {
			h += 1
		}
	}
	return h
}
