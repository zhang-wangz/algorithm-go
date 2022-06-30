package solution

// https://leetcode.cn/problems/range-module/

type SegmentNode struct {
	Ls, Rs *SegmentNode
	// add 代表懒标记
	Add int
	// 代表要统计的值
	Val int
}

func (node *SegmentNode) update(lc int, rc int, l int, r int, v int) {
	le := rc - lc + 1
	if l <= lc && rc <= r {
		node.Add = v
		if v == 1 {
			node.Val = le
		} else {
			node.Val = 0
		}
		return
	}
	node.pushdown(le)
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
	node.pushdown(rc - lc + 1)
	mid, ans := (lc+rc)>>1, 0
	if l <= mid {
		ans = node.Ls.query(lc, mid, l, r)
	}
	if r > mid {
		ans += node.Rs.query(mid+1, rc, l, r)
	}
	return ans
}

func (node *SegmentNode) pushup() {
	node.Val = node.Ls.Val + node.Rs.Val
}

func (node *SegmentNode) pushdown(len int) {
	if node.Ls == nil {
		node.Ls = &SegmentNode{nil, nil, 0, 0}
	}
	if node.Rs == nil {
		node.Rs = &SegmentNode{nil, nil, 0, 0}
	}
	if node.Add == 0 {
		return
	}
	add := node.Add
	if add == -1 {
		node.Ls.Val, node.Rs.Val = 0, 0
	} else {
		node.Ls.Val = len - len/2
		node.Rs.Val = len / 2
	}
	node.Ls.Add, node.Rs.Add = node.Add, node.Add
	node.Add = 0
}

const (
	N1 = 1000000010
)

type RangeModule struct {
	Root *SegmentNode
}

func Constructor1() RangeModule {
	return RangeModule{&SegmentNode{nil, nil, 0, 0}}
}

func (node *RangeModule) AddRange(left int, right int) {
	node.Root.update(1, N-1, left, right-1, 1)
}

func (node *RangeModule) QueryRange(left int, right int) bool {
	return node.Root.query(1, N-1, left, right-1) == right-left
}

func (node *RangeModule) RemoveRange(left int, right int) {
	node.Root.update(1, N-1, left, right-1, -1)
}
