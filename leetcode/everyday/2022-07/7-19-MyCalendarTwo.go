package _022_07

import "math"

type SegmentNode struct {
	Ls, Rs *SegmentNode
	// add 代表懒标记
	Add int
	// 代表要统计的值
	Val int
}

func (node *SegmentNode) update(lc int, rc int, l int, r int, v int) {
	if l <= lc && rc <= r {
		node.Add += v
		node.Val += v
		return
	}
	node.pushdown()
	mid := (lc + rc) >> 1
	if l <= mid {
		node.Ls.update(lc, mid, l, r, v)
	}
	if r > mid {
		node.Rs.update(mid+1, rc, l, r, v)
	}
	node.pushup()
}

func (node *SegmentNode) query(lc int, rc int, l int, r int) int {
	if l <= lc && rc <= r {
		return node.Val
	}
	node.pushdown()
	mid, ans := (lc+rc)>>1, 0
	if l <= mid {
		ans = node.Ls.query(lc, mid, l, r)
	}
	if r > mid {
		right := node.Rs.query(mid+1, rc, l, r)
		if right > ans {
			ans = right
		}
	}
	return ans
}

func (node *SegmentNode) pushup() {
	node.Val = int(math.Max(float64(node.Ls.Val), float64(node.Rs.Val)))
}

func (node *SegmentNode) pushdown() {
	if node.Ls == nil {
		node.Ls = new(SegmentNode)
	}
	if node.Rs == nil {
		node.Rs = new(SegmentNode)
	}
	if node.Add == 0 {
		return
	}
	node.Ls.Val += node.Add
	node.Ls.Add += node.Add
	node.Rs.Val += node.Add
	node.Rs.Add += node.Add
	node.Add = 0
}

const (
	N = 1000000010
)

type MyCalendarTwo struct {
	*SegmentNode
}

func Constructor4() MyCalendarTwo {
	return MyCalendarTwo{&SegmentNode{}}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if this.query(0, N, start, end-1) == 2 {
		return false
	}
	this.update(0, N, start, end-1, 1)
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
