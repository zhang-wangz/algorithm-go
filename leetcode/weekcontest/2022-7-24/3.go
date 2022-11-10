package main

import (
	heap2 "container/heap"
	"sort"
)

type NumberContainers struct {
	idx2num map[int]int
	num2idx map[int]*heap
}

func Constructor() NumberContainers {
	return NumberContainers{idx2num: map[int]int{}, num2idx: map[int]*heap{}}
}

func (this *NumberContainers) Change(index int, number int) {
	this.idx2num[index] = number
	if this.num2idx[number] == nil {
		this.num2idx[number] = &heap{[]int{}}
	}
	h := this.num2idx[number]
	heap2.Push(h, index)
}

func (this *NumberContainers) Find(number int) int {
	h := this.num2idx[number]
	if h == nil {
		return -1
	}
	for h.Len() != 0 {
		idx := heap2.Pop(h).(int)
		if this.idx2num[idx] == number {
			heap2.Push(h, idx)
			return idx
		}
	}
	return -1
}

type heap struct {
	sort.IntSlice
}

func (h *heap) Len() int {
	return len(h.IntSlice)
}
func (h *heap) Push(i interface{}) {
	h.IntSlice = append(h.IntSlice, i.(int))
}
func (h *heap) Pop() interface{} {
	top := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return top
}

/**
* Your NumberContainers object will be instantiated and called as such:
* obj := Constructor();
* obj.Change(index,number);
* param_2 := obj.Find(number);
 */
