package main

import "algorithm-go/base/sort/solution"

func main() {
	nums := []int{9, 2, 5, 4, 8, 2}
	nums = solution.HeapSort(nums)
	for _, val := range nums {
		println(val)
	}

}
