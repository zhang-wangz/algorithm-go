package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, d int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].d = l, r, 1e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 区间赋值
func (t seg) modify(o, l, r, val int) {
	if l <= t[o].l && t[o].r <= r {
		t[o].d = val
		return
	}
	if d := t[o].d; d != 1e9 {
		t[o<<1].d = d
		t[o<<1|1].d = d
		t[o].d = 1e9
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.modify(o<<1, l, r, val)
	}
	if m < r {
		t.modify(o<<1|1, l, r, val)
	}
}

// query
func (t seg) query(o, i int) int {
	if t[o].d != 1e9 || t[o].l == t[o].r {
		return t[o].d
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		return t.query(o<<1, i)
	}
	return t.query(o<<1|1, i)
}

// https://codeforces.com/problemset/problem/292/E

// 输入 n(≤1e5) 和 m (≤1e5)，两个长度都为 n 的数组 a 和 b（元素范围在 [-1e9,1e9] 内，下标从 1 开始）。
// 然后输入 m 个操作：
// 操作 1 形如 1 x y k，表示把 a 的区间 [x,x+k-1] 的元素拷贝到 b 的区间 [y,y+k-1] 上（输入保证下标不越界）。
// 操作 2 形如 2 x，输出 b[x]。

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}
	var op int
	t := make(seg, n*4)
	t.build(1, 1, n)
	for i := 0; i < m; i++ {
		Fscan(in, &op)
		if op == 1 {
			var x, y, k int
			Fscan(in, &x, &y, &k)
			t.modify(1, y, y+k-1, x-y)
		} else {
			var x int
			Fscan(in, &x)
			d := t.query(1, x)
			if d == 1e9 {
				Fprintln(out, b[x])
			} else {
				Fprintln(out, a[x+d])
			}
		}
	}
}

func main() {
	open, _ := os.Open("/Users/zhang-wangz/Desktop/algorithm/tea_0x3f/00_cfinput.txt")
	solve(open, os.Stdout)
	defer open.Close()
	// solve(os.Stdin, os.Stdout)
}
