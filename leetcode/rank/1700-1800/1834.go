package main

import (
	"container/heap"
	"sort"
)

// 单线程cpu
func getOrder(tasks [][]int) (ans []int) {
	n := len(tasks)
	idxs := make([][3]int, n)
	for i := range idxs {
		idxs[i] = [3]int{tasks[i][0], tasks[i][1], i}
	}
	// 对入队时间进行排序
	sort.Slice(idxs, func(i, j int) bool {
		return idxs[i][0] < idxs[j][0]
	})
	h := &hp{}
	time := 1
	idx := 0
	for len(ans) != n {
		for idx < n && idxs[idx][0] <= time {
			h.push(idxs[idx][:])
			idx++
		}
		if h.Len() == 0 {
			time = idxs[idx][0]
		} else {
			q := h.pop()
			ans = append(ans, q[2])
			time += q[1]
		}
	}
	return ans
}

type hp [][]int

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	if h[i][1] != h[j][1] {
		return h[i][1] < h[j][1]
	}
	return h[i][2] < h[j][2]
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v []int)       { heap.Push(h, v) }
func (h *hp) pop() []int         { return heap.Pop(h).([]int) } // 稍微封装一下，方便使用
