package _1_78

import "math/rand"

// 前缀和加二分
// 其实可以准备一个队列，有多少个w权重就放多少个，但看范围，空间很可能超

type Solution struct {
	que []int
}

func Constructor(w []int) Solution {
	q := make([]int, len(w)+1)
	q[0] = 0
	for i := 1; i <= len(w); i++ {
		q[i] = q[i-1] + w[i-1]
	}
	return Solution{
		que: q,
	}
}

func (this *Solution) PickIndex() int {
	// 求小于等于x的第一个数 // 即x的右边界
	x := rand.Intn(this.que[len(this.que)-1])
	n := len(this.que)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)>>1
		if this.que[mid] <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}
