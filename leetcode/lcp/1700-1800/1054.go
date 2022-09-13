package main

import "container/heap"

// 1054
// 贪心 使用出现频率最大的数，如果放入的时候最大的数与最后一个数相等，则从第二大的数据里面取
func rearrangeBarcodes(barcodes []int) (ans []int) {
	cnt := map[int]int{}
	for _, v := range barcodes {
		cnt[v]++
	}
	hp := &myHeap{}
	for k, v := range cnt {
		heap.Push(hp, []int{k, v})
	}
	for hp.Len() > 0 {
		top := heap.Pop(hp).([]int)
		if len(ans) > 0 && ans[len(ans)-1] == top[0] {
			top2 := heap.Pop(hp).([]int)
			ans = append(ans, top2[0])
			if top2[1] > 1 {
				heap.Push(hp, []int{top2[0], top2[1] - 1})
			}
			heap.Push(hp, []int{top[0], top[1]})
		} else {
			ans = append(ans, top[0])
			if top[1] > 1 {
				heap.Push(hp, []int{top[0], top[1] - 1})
			}
		}
	}
	return
}

type myHeap [][]int

func (h myHeap) Len() int {
	return len(h)
}
func (h myHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Pop() interface{} {
	old := *h
	res := old[len(old)-1]
	*h = old[:len(old)-1]
	return res

}
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
