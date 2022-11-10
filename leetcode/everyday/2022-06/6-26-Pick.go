package _022_06

import "math/rand"

// https://leetcode.cn/problems/random-pick-with-blacklist/
type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor2(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m
	black := map[int]bool{}
	b2w := make(map[int]int, 0)
	for _, v := range blacklist {
		if v >= bound {
			black[v] = true
		}
	}
	w := n - m
	for _, v := range blacklist {
		if v < bound {
			for black[w] {
				w++
			}
			b2w[v] = w
			w++
		}
	}
	return Solution{bound: bound, b2w: b2w}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.bound)
	if this.b2w[x] > 0 {
		return this.b2w[x]
	}
	return x
}
