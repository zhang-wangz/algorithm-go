package _1_60

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
		if right >= ans {
			ans = right
		}
	}
	return ans
}

func (node *SegmentNode) pushup() {
	if node.Ls.Val > node.Rs.Val {
		node.Val = node.Ls.Val
	} else {
		node.Val = node.Rs.Val
	}
}

func (node *SegmentNode) pushdown() {
	if node.Ls == nil {
		node.Ls = &SegmentNode{nil, nil, 0, 0}
	}
	if node.Rs == nil {
		node.Rs = &SegmentNode{nil, nil, 0, 0}
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
	N = 1000000000
)

type MyCalendar struct {
	*SegmentNode
}

func Constructor() MyCalendar {
	return MyCalendar{&SegmentNode{Rs: nil, Ls: nil, Add: 0, Val: 0}}
}

func (node *MyCalendar) Book(start int, end int) bool {
	if node.query(0, N, start, end-1) == 0 {
		node.update(0, N, start, end-1, 1)
		return true
	} else {
		return false
	}
}
