package main

func minEatingSpeed(piles []int, h int) int {
	m := piles[0]
	for _, v := range piles {
		if v > m {
			m = v
		}
	}
	l, r := 1, m+1
	for l < r {
		mid := l + (r-l)>>1
		if getH(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func getH(piles []int, k int) int {
	h := 0
	for _, v := range piles {
		if v%k != 0 {
			h += v/k + 1
		} else {
			h += v / k
		}
	}
	return h
}
