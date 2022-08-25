package main

import "sort"

func findClosestElements(arr []int, k, x int) []int {
	// 找到最>=x的位置
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}

func findClosestElements2(arr []int, k, x int) []int {
	// 稳定排序，在绝对值相同的情况下，保证更小的数排在前面
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
