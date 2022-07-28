package main

import "sort"

func arrayRankTransform(arr []int) []int {
	a := append([]int{}, arr...)
	sort.Ints(arr)
	n2i := map[int]int{}
	idx := 1
	for _, n := range arr {
		if _, ok := n2i[n]; !ok {
			n2i[n] = idx
			idx++
		}
	}
	for i, n := range a {
		a[i] = n2i[n]
	}
	return a
}
