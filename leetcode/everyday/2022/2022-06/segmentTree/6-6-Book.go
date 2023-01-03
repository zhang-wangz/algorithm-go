package segmentTree

// https://leetcode.cn/problems/my-calendar-iii/
type MyCalendarThree1 map[int]pair

// tree记录最大值
type pair struct {
	tree, lazy int
}

func Constructor1() MyCalendarThree1 {
	return MyCalendarThree1{}
}

func (t MyCalendarThree1) update(start, end int, l, r int, idx int) {
	if r < start || l > end {
		return
	}
	if start <= l && r <= end {
		tmp := t[idx]
		tmp.tree++
		tmp.lazy++
		t[idx] = tmp
	} else {
		ln, rn := 2*idx, 2*idx+1
		mid := l + (r-l)>>1
		t.update(start, end, l, mid, ln)
		t.update(start, end, mid+1, r, rn)
		tmp := t[idx]
		tmp.tree = tmp.lazy + max(t[ln].tree, t[rn].tree)
		t[idx] = tmp
	}
}
func (t *MyCalendarThree1) Book1(start int, end int) int {
	t.update(start, end-1, 0, 1e9, 1)
	return (*t)[1].tree
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
