package binarysearch

// https://leetcode.cn/problems/koko-eating-bananas/
// 珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
// 珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。
// 每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	l := 1
	r := max
	for l+1 < r {
		mid := l + (r-l)>>1
		if getTm(piles, mid) <= h {
			r = mid
		} else {
			l = mid
		}
	}
	if getTm(piles, l) <= h {
		return l
	}
	if getTm(piles, r) <= h {
		return r
	}
	return l
}

func getTm(piles []int, speed int) int {
	time := 0
	for _, pile := range piles {
		// 向上取整
		time += (pile + speed - 1) / speed
	}
	return time
}
