package solution

import (
	"math"
	"math/rand"
)

// https://leetcode.cn/problems/random-point-in-non-overlapping-rectangles/
// 水塘抽样
type Solution2 struct {
	edge [][]int
}

func Constructor2(rects [][]int) Solution2 {
	return Solution2{edge: rects}
}

func (this *Solution2) Pick() []int {
	idx := 0
	curSum := 0
	for i, e := range this.edge {
		x1, x2 := e[0], e[2]
		y1, y2 := e[1], e[3]
		cur := (x2 - x1 + 1) * (y2 - y1 + 1)
		curSum += cur
		// 当前方块抽到的概率
		if rand.Intn(curSum+1) < cur {
			idx = i
		}
	}
	x2, x1 := this.edge[idx][2], this.edge[idx][0]
	y2, y1 := this.edge[idx][3], this.edge[idx][1]
	x := int(math.Floor(rand.Float64()*float64(x2-x1+1) + float64(x1)))
	y := int(math.Floor(rand.Float64()*float64(y2-y1+1) + float64(y1)))
	return []int{x, y}
}

// 前缀和+二分
type Solution3 struct {
	edge [][]int
	cur  []int
}

func Constructor3(rects [][]int) Solution3 {
	s := Solution3{}
	s.edge = rects
	s.cur = []int{0}
	for _, e := range rects {
		x1, x2 := e[0], e[2]
		y1, y2 := e[1], e[3]
		top := s.cur[len(s.cur)-1]
		s.cur = append(s.cur, top+(x2-x1+1)*(y2-y1+1))
	}
	return s
}

func (this *Solution3) Pick2() []int {
	// [1-top]
	idx := 0
	n := rand.Intn(this.cur[len(this.cur)-1]) + 1
	l := 0
	r := len(this.cur) - 1
	for l+1 < r {
		mid := l + (r-l)>>1
		if this.cur[mid] < n {
			l = mid
		} else {
			r = mid
		}
	}

	if this.cur[l] <= n {
		idx = l
	} else if this.cur[r] <= n {
		idx = r
	}

	x2, x1 := this.edge[idx][2], this.edge[idx][0]
	y2, y1 := this.edge[idx][3], this.edge[idx][1]
	x := int(math.Floor(rand.Float64()*float64(x2-x1+1) + float64(x1)))
	y := int(math.Floor(rand.Float64()*float64(y2-y1+1) + float64(y1)))
	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
